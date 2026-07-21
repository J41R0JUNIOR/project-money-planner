package dto

import (
	root "money-manager/internal/domain"
	domain "money-manager/internal/domain/planner"
)

type CreateEventRequestDTO struct {
	UserId             string                    `json:"user_id"`
	AccountId          string                    `json:"account_id"`
	CategoryId         string                    `json:"category_id"`
	Name               string                    `json:"name"`
	Description        string                    `json:"description"`
	PlannedEventStatus domain.PlannedEventStatus `json:"planned_event_status"`
	Amount             root.Money                `json:"amount"`
}
