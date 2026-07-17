package domain

type Money struct {
	Amount int64  // Amount in cents
	Currency Currency
}

type Currency string

const (
	USD Currency = "USD"
	EUR Currency = "EUR"
	GBP Currency = "GBP"
	JPY Currency = "JPY"
)