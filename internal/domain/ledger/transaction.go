package domain

import (
	root "money-manager/internal/domain"
	"time"
)

type Transaction struct {
	Id        string
	UserId    string
	EventId   string
	AccountId string

	Amount root.Money
	Date   time.Time
	Status TransactionStatus
}

type TransactionStatus string

const (
	Pending   TransactionStatus = "PENDING"
	Completed TransactionStatus = "COMPLETED"
	Failed    TransactionStatus = "FAILED"
)
