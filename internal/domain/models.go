package domain

import "time"

type User struct {
	ID       int
	Username string
	Email    string
}

type Projection struct {
	Date    time.Time
	Balance float64
}

type Transaction struct {
	ID          int
	Type        TypeTransaction
	Amount      float64
	Date        string
	Description string
	IsRecurring bool
}

type TypeTransaction string

const (
	INCOME  TypeTransaction = "INCOME"
	EXPENSE TypeTransaction = "EXPENSE"
)