package domain

type Category struct {
	Id     string
	UserId string

	Name  string
	Icon  string
	Color string

	Type FinancialType
}

type FinancialType string

const (
	Income     FinancialType = "INCOME"
	Expense    FinancialType = "EXPENSE"
	Transfer   FinancialType = "TRANSFER"
	Investment FinancialType = "INVESTMENT"
)
