package user

import "github.com/J41R0JUNIOR/project-money-planner/internal/domain"

type UserService struct{}

func NewService() *UserService {
	return &UserService{}
}

func (s *UserService) Create(user *domain.User) error {
	_ = user
	return nil
}

func (s *UserService) GetByID(id int) (*domain.User, error) {
	_ = id
	return &domain.User{}, nil
}
