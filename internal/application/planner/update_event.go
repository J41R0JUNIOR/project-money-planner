package planner

import (
	"context"

	domain "money-manager/internal/domain/planner"
	"money-manager/internal/repository"
)

type UpdateEventUseCase struct {
	plannerRepository repository.PlannerRepository
}

func NewUpdateEventUseCase(
	plannerRepository repository.PlannerRepository,
) *UpdateEventUseCase {
	return &UpdateEventUseCase{
		plannerRepository: plannerRepository,
	}
}

func (uc *UpdateEventUseCase) Execute(
	ctx context.Context,
	event domain.PlannedEvent,
) error {
	error := uc.plannerRepository.UpdateEvent(event, ctx)

	if error != nil {
		return error
	}

	return nil
}