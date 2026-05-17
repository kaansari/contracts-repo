package security

import "context"

type authUserContextKey struct{}

type AuthenticatedUser struct {
	ID      string
	Name    string
	Email   string
	Company string
	Role    string
}

type ValidatedUser = AuthenticatedUser

func WithAuthenticatedUser(ctx context.Context, user *AuthenticatedUser) context.Context {
	if user == nil {
		return ctx
	}
	return context.WithValue(ctx, authUserContextKey{}, user)
}

func AuthenticatedUserFromContext(ctx context.Context) (*AuthenticatedUser, bool) {
	user, ok := ctx.Value(authUserContextKey{}).(*AuthenticatedUser)
	return user, ok && user != nil
}
