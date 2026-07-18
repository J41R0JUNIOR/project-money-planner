package auth

import (
	"context"
	provider "money-manager/internal/provider"
)

type RefreshUseCase struct {
	authProvider provider.AuthProvider
}

func NewRefreshUseCase(authProvider provider.AuthProvider) *RefreshUseCase {
	return &RefreshUseCase{
		authProvider: authProvider,
	}
}

func (uc *RefreshUseCase) Execute(
	ctx context.Context,
	refreshToken string,
) (AuthResult, error) {
	auth, err := uc.authProvider.Refresh(
		ctx,
		refreshToken,
	)

	if err != nil {
		return AuthResult{}, err
	}

	return AuthResult{
		AccessToken:  auth.AccessToken,
		IdToken:      auth.IdToken,
		RefreshToken: auth.RefreshToken,
	}, nil
}