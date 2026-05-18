package security

import (
	"context"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type fakePermissionChecker struct {
	allowed bool
	err     error
	role    string
	method  string
}

func (c *fakePermissionChecker) IsAllowed(ctx context.Context, role string, grpcMethod string) (bool, error) {
	c.role = role
	c.method = grpcMethod
	return c.allowed, c.err
}

func (c *fakePermissionChecker) Refresh(ctx context.Context) error {
	return c.err
}

func TestRBACPublicMethodBypassesChecker(t *testing.T) {
	checker := &fakePermissionChecker{}
	interceptor := NewRBACInterceptor(checker, []string{"/auth.Auth/Auth"})

	_, err := interceptor.Unary()(context.Background(), "request", &grpc.UnaryServerInfo{FullMethod: "/auth.Auth/Auth"}, func(ctx context.Context, req any) (any, error) {
		return "ok", nil
	})
	if err != nil {
		t.Fatalf("expected public method to pass, got %v", err)
	}
	if checker.method != "" {
		t.Fatalf("expected checker not to be called, got method %q", checker.method)
	}
}

func TestRBACRejectsUnauthenticatedProtectedMethod(t *testing.T) {
	interceptor := NewRBACInterceptor(&fakePermissionChecker{allowed: true}, nil)

	_, err := interceptor.Unary()(context.Background(), "request", &grpc.UnaryServerInfo{FullMethod: "/customer.CustomerService/ListCustomers"}, nil)
	if status.Code(err) != codes.Unauthenticated {
		t.Fatalf("expected Unauthenticated, got %v", err)
	}
}

func TestRBACRejectsDeniedRole(t *testing.T) {
	interceptor := NewRBACInterceptor(&fakePermissionChecker{}, nil)
	ctx := WithAuthenticatedUser(context.Background(), &AuthenticatedUser{ID: "user-1", Role: "customer"})

	_, err := interceptor.Unary()(ctx, "request", &grpc.UnaryServerInfo{FullMethod: "/customer.CustomerService/ListCustomers"}, nil)
	if status.Code(err) != codes.PermissionDenied {
		t.Fatalf("expected PermissionDenied, got %v", err)
	}
}

func TestRBACAllowsPermittedRole(t *testing.T) {
	checker := &fakePermissionChecker{allowed: true}
	interceptor := NewRBACInterceptor(checker, nil)
	ctx := WithAuthenticatedUser(context.Background(), &AuthenticatedUser{ID: "user-1", Role: "agent"})

	_, err := interceptor.Unary()(ctx, "request", &grpc.UnaryServerInfo{FullMethod: "/customer.CustomerService/ListCustomers"}, func(ctx context.Context, req any) (any, error) {
		return "ok", nil
	})
	if err != nil {
		t.Fatalf("expected allowed role to pass, got %v", err)
	}
	if checker.role != "agent" || checker.method != "/customer.CustomerService/ListCustomers" {
		t.Fatalf("unexpected checker input: %#v", checker)
	}
}
