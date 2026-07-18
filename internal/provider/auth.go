package provider

import (
	"context"
	domain "money-manager/internal/domain/auth"
)

type AuthProvider interface {
	SignUp(ctx context.Context, userName string, password string, userEmail string) (string, error)
	SignIn(ctx context.Context, email string, password string) (domain.Auth, error)
	Refresh(ctx context.Context, refreshToken string) (domain.Auth, error)
}