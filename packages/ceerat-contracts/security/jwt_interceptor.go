package security

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"log/slog"
	"strings"

	authpb "github.com/kaansari/ceerat-platform/packages/ceerat-contracts/proto/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type TokenValidator interface {
	ValidateToken(ctx context.Context, token string) (*ValidatedUser, error)
}

type JWTInterceptor struct {
	validator TokenValidator
	allowlist map[string]bool
	logger    *slog.Logger
}

func NewJWTInterceptor(validator TokenValidator, publicMethods []string) *JWTInterceptor {
	allowlist := make(map[string]bool, len(publicMethods))
	for _, method := range publicMethods {
		method = strings.TrimSpace(method)
		if method != "" {
			allowlist[method] = true
		}
	}
	return &JWTInterceptor{
		validator: validator,
		allowlist: allowlist,
		logger:    slog.Default(),
	}
}

func (i *JWTInterceptor) WithLogger(logger *slog.Logger) *JWTInterceptor {
	if logger != nil {
		i.logger = logger
	}
	return i
}

func (i *JWTInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		if i.isPublic(info.FullMethod) {
			return handler(ctx, req)
		}

		authCtx, err := i.authenticate(ctx, info.FullMethod)
		if err != nil {
			return nil, err
		}
		return handler(authCtx, req)
	}
}

func (i *JWTInterceptor) Stream() grpc.StreamServerInterceptor {
	return func(srv any, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		if i.isPublic(info.FullMethod) {
			return handler(srv, stream)
		}

		authCtx, err := i.authenticate(stream.Context(), info.FullMethod)
		if err != nil {
			return err
		}
		return handler(srv, &authenticatedServerStream{ServerStream: stream, ctx: authCtx})
	}
}

func (i *JWTInterceptor) isPublic(method string) bool {
	return i.allowlist[method]
}

func (i *JWTInterceptor) authenticate(ctx context.Context, method string) (context.Context, error) {
	token, ok := bearerTokenFromContext(ctx)
	if !ok {
		i.log(ctx, "grpc.auth.missing_token", "method", method)
		return nil, status.Error(codes.Unauthenticated, "authentication required")
	}
	if i.validator == nil {
		i.log(ctx, "grpc.auth.invalid_token", "method", method)
		return nil, status.Error(codes.Unauthenticated, "invalid token")
	}

	user, err := i.validator.ValidateToken(ctx, token)
	if err != nil || user == nil || strings.TrimSpace(user.ID) == "" {
		i.log(ctx, "grpc.auth.invalid_token", "method", method)
		return nil, status.Error(codes.Unauthenticated, "invalid token")
	}

	i.log(ctx, "grpc.auth.validated", "method", method, "user_id", user.ID)
	return WithAuthenticatedUser(ctx, user), nil
}

func bearerTokenFromContext(ctx context.Context) (string, bool) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", false
	}

	for _, key := range []string{"authorization", "Authorization"} {
		for _, value := range md.Get(key) {
			token, ok := tokenFromAuthorization(value)
			if ok {
				return token, true
			}
		}
	}
	for _, value := range md.Get("x-auth-token") {
		if token := strings.TrimSpace(value); token != "" {
			return token, true
		}
	}
	return "", false
}

func tokenFromAuthorization(value string) (string, bool) {
	value = strings.TrimSpace(value)
	if len(value) < len("Bearer ")+1 || !strings.EqualFold(value[:len("Bearer")], "Bearer") {
		return "", false
	}
	token := strings.TrimSpace(value[len("Bearer"):])
	return token, token != ""
}

func (i *JWTInterceptor) log(ctx context.Context, event string, args ...any) {
	if i.logger != nil {
		i.logger.InfoContext(ctx, event, args...)
	}
}

type authenticatedServerStream struct {
	grpc.ServerStream
	ctx context.Context
}

func (s *authenticatedServerStream) Context() context.Context {
	return s.ctx
}

type UserServiceTokenValidator struct {
	client authpb.AuthClient
}

func NewUserServiceTokenValidator(client authpb.AuthClient) *UserServiceTokenValidator {
	return &UserServiceTokenValidator{client: client}
}

func (v *UserServiceTokenValidator) ValidateToken(ctx context.Context, token string) (*ValidatedUser, error) {
	if v == nil || v.client == nil {
		return nil, errors.New("auth client is required")
	}
	resp, err := v.client.ValidateToken(ctx, &authpb.Token{Token: token})
	if err != nil {
		return nil, err
	}
	if !resp.GetValid() {
		return nil, errors.New("invalid token")
	}
	return userFromJWT(token)
}

func userFromJWT(token string) (*ValidatedUser, error) {
	parts := strings.Split(token, ".")
	if len(parts) < 2 {
		return nil, errors.New("invalid jwt format")
	}
	payload, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return nil, err
	}
	var claims struct {
		User struct {
			ID      string `json:"id"`
			Name    string `json:"name"`
			Email   string `json:"email"`
			Company string `json:"company"`
			Role    string `json:"role"`
			Status  string `json:"status"`
		} `json:"user"`
	}
	if err := json.Unmarshal(payload, &claims); err != nil {
		return nil, err
	}
	if claims.User.ID == "" {
		return nil, errors.New("jwt does not contain user.id")
	}
	return &ValidatedUser{
		ID:      claims.User.ID,
		Name:    claims.User.Name,
		Email:   claims.User.Email,
		Company: claims.User.Company,
		Role:    claims.User.Role,
		Status:  claims.User.Status,
	}, nil
}
