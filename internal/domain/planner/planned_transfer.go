package domain

import (
	"time"

	root "money-manager/internal/domain"
)

type PlannedTransfer struct {
    Id string

    FromAccountId string
    ToAccountId   string

    Amount root.Money

    StartDate time.Time
    Recurrence *Recurrence
}