package repository

import (
	domain "money-manager/internal/domain/ledger"
	"context"
)

type LedgerRepository interface {
	Save(ledger domain.Transaction, ctx context.Context) error
}