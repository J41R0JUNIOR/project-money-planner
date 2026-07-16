package service

import (
	"context"

	"money-manager/internal/api/auth/dto"
	"money-manager/internal/repository"
	"money-manager/internal/domain"
	"money-manager/internal/provider"
)

type AuthService struct {
	authProvider    provider.AuthProvider
	userRepository repository.UserRepository
}

func NewAuthService(authProvider provider.AuthProvider,userRepository repository.UserRepository) *AuthService {
	return &AuthService{
		userRepository: userRepository,
		authProvider:   authProvider,
	}
}

func (s *AuthService) SignUp(
	ctx context.Context,
	request auth_dto.SignUpRequestDTO,
) error {
	

	// Criar conta com cognito primeiro
	s.userRepository.Save( 
			domain.User{
			Id: "generated-id",
			Email: request.Email,
			Name: request.Name,
		},
	)
	
	return nil
}

func (s *AuthService) SignIn(
	ctx context.Context,
	request auth_dto.SignInRequestDTO,
) (string, error) {

	// validar email

	// verificar se usuário existe

	// comparar senha

	// gerar token	
	return "token", nil
}