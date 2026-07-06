package repository

import "github.com/J41R0JUNIOR/project-money-planner/internal/domain"

type UserRepository interface {
	CreateUser(user *domain.User) error
	GetUserByID(id int) (*domain.User, error)
	DeleteUser(id int) error
}
