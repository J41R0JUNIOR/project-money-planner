package domain


// Amount in cents 
type Money struct {
	Amount int64  `json:"amount"`
	Currency Currency `json:"currency"`
}

type Currency string

const (
	USD Currency = "USD"
	EUR Currency = "EUR"
	GBP Currency = "GBP"
	JPY Currency = "JPY"
)