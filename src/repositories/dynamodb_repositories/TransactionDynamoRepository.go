package dynamodb_repositories

import (
	"github.com/J41R0JUNIOR/project-money-planner/src/domain"
	"github.com/J41R0JUNIOR/project-money-planner/src/repositories"
)

type TransactionDynamoRepository struct {
	// DynamoDB client and other necessary fields
}

// CreateTransaction implements [repositories.TransactionRepository].
func (t *TransactionDynamoRepository) CreateTransaction(tx *domain.Transaction) error {
	panic("unimplemented")
}

// DeleteTransaction implements [repositories.TransactionRepository].
func (t *TransactionDynamoRepository) DeleteTransaction(id int) error {
	panic("unimplemented")
}

// GetTransactionByID implements [repositories.TransactionRepository].
func (t *TransactionDynamoRepository) GetTransactionByID(id int) (*domain.Transaction, error) {
	panic("unimplemented")
}

// ListTransactions implements [repositories.TransactionRepository].
func (t *TransactionDynamoRepository) ListTransactions() ([]*domain.Transaction, error) {
	panic("unimplemented")
}

// Ensure TransactionDynamoRepository implements the TransactionRepository interface
var _ repositories.TransactionRepository = (*TransactionDynamoRepository)(nil)
