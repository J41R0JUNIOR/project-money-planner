package repository

import (
	domain "money-manager/internal/domain/account"
	"context"
)

type AccountRepository interface {
	// GetById(accountID string, ctx context.Context) (domain.Account, error)
	// Get(ctx context.Context) ([]domain.Account, error)
	Save(account domain.Account, ctx context.Context) error
	Delete(accountID string, ctx context.Context) error
}