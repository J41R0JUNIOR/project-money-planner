package repository

import (
	domain "money-manager/internal/domain/user"
)
type UserRepository interface {
	Save(user domain.User) error
	Delete(id string) error
}