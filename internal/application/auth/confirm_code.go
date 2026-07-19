package auth

import (
	"context"
	provider "money-manager/internal/provider"
)

type ConfirmCodeUseCase struct {
	authProvider provider.AuthProvider
}

func NewConfirmCodeUseCase(authProvider provider.AuthProvider) *ConfirmCodeUseCase {
	return &ConfirmCodeUseCase{
		authProvider: authProvider,
	}
}

func (uc *ConfirmCodeUseCase) Execute(
	ctx context.Context,
	email string,
	code string,
) (string, error) {
	err := uc.authProvider.ConfirmCode(
		ctx,
		email,
		code,
	)
	if err != nil {
		return "Error to confirm code", err
	}
	return "Code confirmed successfully", nil
}