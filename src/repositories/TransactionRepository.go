package repositories

import "github.com/J41R0JUNIOR/project-money-planner/src/domain"

type TransactionRepository interface {
	createTransaction(tx *domain.Transaction) error
	getTransactionByID(id int) (*domain.Transaction, error)
	listTransactions() ([]*domain.Transaction, error)
	deleteTransaction(id int) error
}
