package domain

import (
	root "money-manager/internal/domain"
	"time"
)

type PlannedEvent struct {
	Id     string
	UserId string

	AccountId  string
	CategoryId string

	Name        string
	Description string

	Status PlannedEventStatus

	Amount root.Money

	StartDate time.Time

	Recurrence *Recurrence
}

type PlannedEventStatus string

const (
	Active   PlannedEventStatus = "ACTIVE"
	Paused   PlannedEventStatus = "PAUSED"
	Archived PlannedEventStatus = "ARCHIVED"
	Deleted  PlannedEventStatus = "DELETED"
)
