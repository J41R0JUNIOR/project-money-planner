package repository

import (
	domain "money-manager/internal/domain/user"
	"context"
)
type UserRepository interface {
	Save(user domain.User, ctx context.Context) error
}