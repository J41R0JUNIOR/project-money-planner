package repository

import (
	domain "money-manager/internal/domain/planner"
	"context"
)

type PlannerRepository interface {
	SaveCategory(category domain.Category, ctx context.Context) error
	DeleteCategory(categoryID string, ctx context.Context) error
	
	SaveEvent(event domain.PlannedEvent, ctx context.Context) error
	UpdateEvent(event domain.PlannedEvent, ctx context.Context) error
	DeleteEvent(eventID string, ctx context.Context) error
	
	SaveRecurringEvent(event domain.Recurrence, ctx context.Context) error
	DeleteRecurringEvent(eventID string, ctx context.Context) error
	
	SavePlannedTransfer(transfer domain.PlannedTransfer, ctx context.Context) error
	DeletePlannedTransfer(transferID string, ctx context.Context) error
}