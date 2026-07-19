package provider

import (
	"context"
	domain "money-manager/internal/domain/auth"
)

type AuthProvider interface {
	SignUp(ctx context.Context, userEmail string, password string) (string, error)
	SignIn(ctx context.Context, email string, password string) (domain.Auth, error)
	ConfirmCode(ctx context.Context, email string, code string) error
	Refresh(ctx context.Context, refreshToken string) (domain.Auth, error)
}