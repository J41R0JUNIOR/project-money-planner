package repositories

import "github.com/J41R0JUNIOR/project-money-planner/src/domain"

type UserRepository interface {
	CreateUser(user *domain.User) error
	GetUserByID(id int) (*domain.User, error)
	DeleteUser(id int) error
}