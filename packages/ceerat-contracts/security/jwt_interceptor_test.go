package security

import (
	"context"
	"errors"
	"io"
	"log/slog"
	"testing"

	authpb "github.com/kaansari/ceerat-platform/packages/ceerat-contracts/proto/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type fakeValidator struct {
	user *ValidatedUser
	err  error
}

func (v fakeValidator) ValidateToken(ctx context.Context, token string) (*ValidatedUser, error) {
	if token != "valid-token" {
		return nil, errors.New("invalid token")
	}
	if v.err != nil {
		return nil, v.err
	}
	return v.user, nil
}

type fakeAuthClient struct {
	resp *authpb.Token
	err  error
}

func (c fakeAuthClient) Create(ctx context.Context, in *authpb.User, opts ...grpc.CallOption) (*authpb.Response, error) {
	return nil, errors.New("not implemented")
}

func (c fakeAuthClient) RegisterCustomer(ctx context.Context, in *authpb.RegisterCustomerRequest, opts ...grpc.CallOption) (*authpb.Response, error) {
	return nil, errors.New("not implemented")
}

func (c fakeAuthClient) Get(ctx context.Context, in *authpb.User, opts ...grpc.CallOption) (*authpb.Response, error) {
	return nil, errors.New("not implemented")
}

func (c fakeAuthClient) GetAll(ctx context.Context, in *authpb.Request, opts ...grpc.CallOption) (*authpb.Response, error) {
	return nil, errors.New("not implemented")
}

func (c fakeAuthClient) Auth(ctx context.Context, in *authpb.User, opts ...grpc.CallOption) (*authpb.Token, error) {
	return nil, errors.New("not implemented")
}

func (c fakeAuthClient) ValidateToken(ctx context.Context, in *authpb.Token, opts ...grpc.CallOption) (*authpb.Token, error) {
	if c.err != nil {
		return nil, c.err
	}
	return c.resp, nil
}

func (c fakeAuthClient) UpdateProfile(ctx context.Context, in *authpb.User, opts ...grpc.CallOption) (*authpb.Response, error) {
	return nil, errors.New("not implemented")
}

func (c fakeAuthClient) UpdatePassword(ctx context.Context, in *authpb.PasswordUpdate, opts ...grpc.CallOption) (*authpb.Response, error) {
	return nil, errors.New("not implemented")
}

func TestPublicMethodBypassesJWTValidation(t *testing.T) {
	interceptor := NewJWTInterceptor(fakeValidator{err: errors.New("should not be called")}, []string{"/auth.Auth/ValidateToken"}).
		WithLogger(slog.New(slog.NewTextHandler(io.Discard, nil)))

	_, err := interceptor.Unary()(context.Background(), "request", &grpc.UnaryServerInfo{FullMethod: "/auth.Auth/ValidateToken"}, func(ctx context.Context, req any) (any, error) {
		return "ok", nil
	})
	if err != nil {
		t.Fatalf("expected public method to bypass auth, got %v", err)
	}
}

func TestProtectedMethodRejectsMissingToken(t *testing.T) {
	interceptor := NewJWTInterceptor(fakeValidator{}, nil).
		WithLogger(slog.New(slog.NewTextHandler(io.Discard, nil)))

	_, err := interceptor.Unary()(context.Background(), "request", &grpc.UnaryServerInfo{FullMethod: "/customer.CustomerService/CreateCustomer"}, nil)
	if status.Code(err) != codes.Unauthenticated {
		t.Fatalf("expected Unauthenticated, got %v", err)
	}
}

func TestProtectedMethodRejectsInvalidToken(t *testing.T) {
	interceptor := NewJWTInterceptor(fakeValidator{}, nil).
		WithLogger(slog.New(slog.NewTextHandler(io.Discard, nil)))
	ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs("authorization", "Bearer bad-token"))

	_, err := interceptor.Unary()(ctx, "request", &grpc.UnaryServerInfo{FullMethod: "/customer.CustomerService/CreateCustomer"}, nil)
	if status.Code(err) != codes.Unauthenticated {
		t.Fatalf("expected Unauthenticated, got %v", err)
	}
}

func TestProtectedMethodInjectsAuthenticatedUser(t *testing.T) {
	interceptor := NewJWTInterceptor(fakeValidator{user: &ValidatedUser{ID: "user-1", Email: "owner@example.com"}}, nil).
		WithLogger(slog.New(slog.NewTextHandler(io.Discard, nil)))
	ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs("authorization", "Bearer valid-token"))

	_, err := interceptor.Unary()(ctx, "request", &grpc.UnaryServerInfo{FullMethod: "/customer.CustomerService/CreateCustomer"}, func(ctx context.Context, req any) (any, error) {
		user, ok := AuthenticatedUserFromContext(ctx)
		if !ok {
			t.Fatal("expected authenticated user in context")
		}
		if user.ID != "user-1" {
			t.Fatalf("expected user-1, got %q", user.ID)
		}
		return "ok", nil
	})
	if err != nil {
		t.Fatalf("expected valid token to pass, got %v", err)
	}
}

func TestXAuthTokenFallback(t *testing.T) {
	interceptor := NewJWTInterceptor(fakeValidator{user: &ValidatedUser{ID: "user-1"}}, nil).
		WithLogger(slog.New(slog.NewTextHandler(io.Discard, nil)))
	ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs("x-auth-token", "valid-token"))

	_, err := interceptor.Unary()(ctx, "request", &grpc.UnaryServerInfo{FullMethod: "/customer.CustomerService/CreateCustomer"}, func(ctx context.Context, req any) (any, error) {
		return "ok", nil
	})
	if err != nil {
		t.Fatalf("expected x-auth-token fallback to pass, got %v", err)
	}
}

func TestUserServiceTokenValidatorUsesReturnedUser(t *testing.T) {
	validator := NewUserServiceTokenValidator(fakeAuthClient{resp: &authpb.Token{
		Valid: true,
		User: &authpb.User{
			Id:      "user-1",
			Name:    "Current User",
			Email:   "current@example.com",
			Company: "Ceerat",
			Role:    "agent",
			Status:  "active",
		},
	}})

	user, err := validator.ValidateToken(context.Background(), "not-a-jwt")
	if err != nil {
		t.Fatal(err)
	}
	if user.ID != "user-1" || user.Email != "current@example.com" || user.Role != "agent" {
		t.Fatalf("expected validated user from auth response, got %#v", user)
	}
}

func TestUserServiceTokenValidatorRejectsMissingReturnedUser(t *testing.T) {
	validator := NewUserServiceTokenValidator(fakeAuthClient{resp: &authpb.Token{Valid: true}})

	_, err := validator.ValidateToken(context.Background(), "valid-token")
	if err == nil {
		t.Fatal("expected missing user id error")
	}
}
