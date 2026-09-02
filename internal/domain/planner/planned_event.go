package domain

import (
	root "money-manager/internal/domain"
	"time"
)

type PlannedEvent struct {
	Id     string `json:"id"`
	UserId string `json:"userId"`
	AccountId  string `json:"accountId"`
	CategoryId string `json:"categoryId"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status PlannedEventStatus `json:"status"`
	Amount root.Money `json:"amount"`
	StartDate time.Time `json:"startDate"`
	Recurrence *Recurrence `json:"recurrence"`
}

type PlannedEventStatus string

const (
	Active   PlannedEventStatus = "ACTIVE"
	Paused   PlannedEventStatus = "PAUSED"
	Archived PlannedEventStatus = "ARCHIVED"
	Deleted  PlannedEventStatus = "DELETED"
)
