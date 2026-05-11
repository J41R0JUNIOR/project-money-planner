package repositories

import "github.com/J41R0JUNIOR/project-money-planner/src/domain"

type TransactionRepository interface {
	CreateTransaction(tx *domain.Transaction) error
	GetTransactionByID(id int) (*domain.Transaction, error)
	ListTransactions() ([]*domain.Transaction, error)
	DeleteTransaction(id int) error
}
