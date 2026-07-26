package planner

import (
	"context"

	domain "money-manager/internal/domain/planner"
	"money-manager/internal/repository"
)

type ReadEventUseCase struct {
	plannerRepository repository.PlannerRepository
}

func NewReadEventUseCase(
	plannerRepository repository.PlannerRepository,
) *ReadEventUseCase {
	return &ReadEventUseCase{
		plannerRepository: plannerRepository,
	}
}

func (uc *ReadEventUseCase) Execute(
	ctx context.Context,
	UserId string,
) ([]domain.PlannedEvent, error) {
	data, error := uc.plannerRepository.ReadEvent(UserId, ctx)
	if error != nil {
		return nil, error
	}

	return data, nil
}
