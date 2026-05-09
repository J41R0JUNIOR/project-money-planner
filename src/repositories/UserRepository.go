package repositories

import "github.com/J41R0JUNIOR/project-money-planner/src/domain"

type UserRepository interface {
	createUser(user *domain.User) error
	getUserByID(id int) (*domain.User, error)
	deleteUser(id int) error
}