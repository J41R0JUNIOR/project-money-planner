package repository

import (
	domain "money-manager/internal/domain/planner"
	"context"
)

type PlannerRepository interface {
	SaveCategory(category domain.Category, ctx context.Context) error
	DeleteCategory(categoryID string, ctx context.Context) error
	
	SaveEvent(event domain.PlannedEvent, ctx context.Context) error
	DeleteEvent(userId string, eventID string, ctx context.Context) error
	ReadEvent(userId string, ctx context.Context) ([]domain.PlannedEvent, error)
	
	SavePlannedTransfer(transfer domain.PlannedTransfer, ctx context.Context) error
	DeletePlannedTransfer(transferID string, ctx context.Context) error
}