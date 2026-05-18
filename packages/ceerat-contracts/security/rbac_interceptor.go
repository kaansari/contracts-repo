package security

import (
	"context"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type PermissionChecker interface {
	IsAllowed(ctx context.Context, role string, grpcMethod string) (bool, error)
	Refresh(ctx context.Context) error
}

type RBACInterceptor struct {
	checker       PermissionChecker
	publicMethods map[string]bool
}

func NewRBACInterceptor(checker PermissionChecker, publicMethods []string) *RBACInterceptor {
	public := make(map[string]bool, len(publicMethods))
	for _, method := range publicMethods {
		method = strings.TrimSpace(method)
		if method != "" {
			public[method] = true
		}
	}
	return &RBACInterceptor{checker: checker, publicMethods: public}
}

func (i *RBACInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		if i.isPublic(info.FullMethod) {
			return handler(ctx, req)
		}

		user, ok := AuthenticatedUserFromContext(ctx)
		if !ok || strings.TrimSpace(user.ID) == "" {
			return nil, status.Error(codes.Unauthenticated, "authentication required")
		}

		role := strings.TrimSpace(user.Role)
		if role == "" {
			return nil, status.Error(codes.PermissionDenied, "role is not allowed to access this service")
		}
		if i.checker == nil {
			return nil, status.Error(codes.PermissionDenied, "role is not allowed to access this service")
		}

		allowed, err := i.checker.IsAllowed(ctx, role, info.FullMethod)
		if err != nil {
			return nil, status.Error(codes.PermissionDenied, "access denied")
		}
		if !allowed {
			return nil, status.Error(codes.PermissionDenied, "role is not allowed to access this service")
		}

		return handler(ctx, req)
	}
}

func (i *RBACInterceptor) Stream() grpc.StreamServerInterceptor {
	return func(srv any, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		if i.isPublic(info.FullMethod) {
			return handler(srv, stream)
		}

		user, ok := AuthenticatedUserFromContext(stream.Context())
		if !ok || strings.TrimSpace(user.ID) == "" {
			return status.Error(codes.Unauthenticated, "authentication required")
		}

		role := strings.TrimSpace(user.Role)
		if role == "" || i.checker == nil {
			return status.Error(codes.PermissionDenied, "role is not allowed to access this service")
		}

		allowed, err := i.checker.IsAllowed(stream.Context(), role, info.FullMethod)
		if err != nil {
			return status.Error(codes.PermissionDenied, "access denied")
		}
		if !allowed {
			return status.Error(codes.PermissionDenied, "role is not allowed to access this service")
		}

		return handler(srv, stream)
	}
}

func (i *RBACInterceptor) isPublic(method string) bool {
	return i != nil && i.publicMethods[method]
}
