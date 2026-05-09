package domain

type Transaction struct {
	ID          int
	Type        TypeTransaction
	Amount      float64
	Date        string  // ISO format: YYYY-MM-DD
	Description string
	IsRecurring bool
}