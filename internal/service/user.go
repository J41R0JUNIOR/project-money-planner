package service

import (
	"context"
	"money-manager/internal/api/user/dto"
)

type UserService struct {
	// repository UserRepository
}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) CreateUser(
	ctx context.Context,
	request dto.CreateUserRequestDTO,
) error {
	// verificar se usuário existe

	// criptografar senha

	// salvar no banco

	return nil
}