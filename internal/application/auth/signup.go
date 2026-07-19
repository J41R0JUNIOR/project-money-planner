package auth

import (
	"context"

	"money-manager/internal/domain/user"
	"money-manager/internal/provider"
	"money-manager/internal/repository"
)

type SignUpUseCase struct {
	authProvider provider.AuthProvider
	userRepository repository.UserRepository
}

func NewSignUpUseCase(
	authProvider provider.AuthProvider,
	userRepository repository.UserRepository,
) *SignUpUseCase {

	return &SignUpUseCase{
		authProvider: authProvider,
		userRepository: userRepository,
	}
}

func (uc *SignUpUseCase) Execute(
	ctx context.Context,
	name string,
	email string,
	password string,
) error {
	userID, err := uc.authProvider.SignUp(
		ctx,
		email,
		password,
	)

	if err != nil {
		return err
	}

	user := domain.User{
		Id: userID,
		Name: name,
		Email: email,
	}

	return uc.userRepository.Save(user, ctx)
}
