package domain

import (
	"time"
	root "money-manager/internal/domain"
)

type Projection struct {
    Date time.Time

    ExpectedIncome  root.Money
    ExpectedExpense root.Money

    ExpectedBalance root.Money
}