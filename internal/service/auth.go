package service

import (
	"context"

	dto "money-manager/internal/api/auth/dto"
)

type AuthService struct {
	// repository UserRepository
}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (s *AuthService) SignUp(
	ctx context.Context,
	request dto.SignUpRequestDTO,
) error {

	// validar email

	// verificar se usuário existe

	// criptografar senha

	// salvar no banco

	return nil
}

func (s *AuthService) SignIn(
	ctx context.Context,
	request dto.SignInRequestDTO,
) (string, error) {

	// validar email

	// verificar se usuário existe

	// comparar senha

	// gerar token	
	return "token", nil
}