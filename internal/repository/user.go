package repository

import (
	"money-manager/internal/domain"
)

type UserRepository interface {
	Save(user domain.User) error
	Delete(id string) error
}