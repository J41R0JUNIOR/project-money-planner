package dto

import (
	root "money-manager/internal/domain"
	domain "money-manager/internal/domain/planner"
)

type UpdateEventRequestDTO struct {
	Id          string
	CategoryId  string
	Name        string
	Description string
	Status      domain.PlannedEventStatus
	Amount      root.Money
}

func (u UpdateEventRequestDTO) ToDomain() domain.PlannedEvent {
	return domain.PlannedEvent{
		Id:          u.Id,
		CategoryId:  u.CategoryId,
		Name:        u.Name,
		Description: u.Description,
		Status:      u.Status,
		Amount:      u.Amount,
	}
}
