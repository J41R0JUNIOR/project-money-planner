package auth
import (
	"context"
	provider "money-manager/internal/provider"
)

type SignInUseCase struct {
	authProvider provider.AuthProvider
}

func NewSignInUseCase(
	authProvider provider.AuthProvider,
) *SignInUseCase {

	return &SignInUseCase{
		authProvider: authProvider,
	}
}

func (uc *SignInUseCase) Execute(
	ctx context.Context,
	email string,
	password string,
) (SignInResult, error) {


	auth, err := uc.authProvider.SignIn(
		ctx,
		email,
		password,
	)


	if err != nil {
		return SignInResult{}, err
	}


	return SignInResult{
		AccessToken: auth.AccessToken,
		IdToken: auth.IdToken,
		RefreshToken: auth.RefreshToken,
	}, nil
}