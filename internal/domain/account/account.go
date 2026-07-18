package domain

import (
    root "money-manager/internal/domain"
)

type Account struct {
    Id string
    UserId string

    Name string
    Type AccountType
    Balance root.Money
}

type AccountType string

const (
    Checking AccountType = "CHECKING"
    Savings  AccountType = "SAVINGS"
    Credit   AccountType = "CREDIT"
    Investment AccountType = "INVESTMENT"
)