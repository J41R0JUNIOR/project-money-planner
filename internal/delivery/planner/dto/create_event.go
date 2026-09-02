package dto

import (
	root "money-manager/internal/domain"
	domain "money-manager/internal/domain/planner"
)

type CreateEventRequestDTO struct {
	AccountId          string                    `json:"account_id"`
	CategoryId         string                    `json:"category_id"`
	Name               string                    `json:"name"`
	Description        string                    `json:"description"`
	Status             domain.PlannedEventStatus `json:"status"`
	Amount             root.Money                `json:"amount"`
	StartDate          string                    `json:"start_date"`
	Recurrence         *domain.Recurrence        `json:"recurrence,omitempty"`
}
 