package planner

import (
	"context"

	"money-manager/internal/repository"
)

type DeleteEventUseCase struct {
	plannerRepository repository.PlannerRepository
}

func NewDeleteEventUseCase(
	plannerRepository repository.PlannerRepository,
) *DeleteEventUseCase {
	return &DeleteEventUseCase{
		plannerRepository: plannerRepository,
	}
}

func (uc *DeleteEventUseCase) Execute(
	ctx context.Context,
	eventId string,
) error {
	error := uc.plannerRepository.DeleteEvent(eventId, ctx)

	if error != nil {
		return error
	}

	return nil
}