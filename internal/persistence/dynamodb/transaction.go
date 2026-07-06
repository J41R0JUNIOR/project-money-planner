package dynamodb

import (
	"github.com/J41R0JUNIOR/project-money-planner/internal/domain"
)

type TransactionRepository struct {
	client *Client
}

func NewTransactionRepository(client *Client) *TransactionRepository {
	return &TransactionRepository{client: client}
}

func (r *TransactionRepository) Create(tx *domain.Transaction) error {
	_ = tx
	return nil
}
