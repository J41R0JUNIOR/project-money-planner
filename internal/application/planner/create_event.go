package planner

import (
	"context"

	root "money-manager/internal/domain"
	domain "money-manager/internal/domain/planner"
	"money-manager/internal/repository"
)

type CreateEventUseCase struct {
	plannerRepository repository.PlannerRepository
}

func NewCreateEventUseCase(
	plannerRepository repository.PlannerRepository,
) *CreateEventUseCase {
	return &CreateEventUseCase{
		plannerRepository: plannerRepository,
	}
}

func (uc *CreateEventUseCase) Execute(
	ctx context.Context,
	UserId string,
	AccountId string,
	CategoryId string,
	Name string,
	Description string,
	Status domain.PlannedEventStatus,
	Amount root.Money,
) error {
	var newEvent = domain.PlannedEvent{
		Id:          "",
		UserId:      UserId,
		AccountId:   AccountId,
		CategoryId:  CategoryId,
		Name:        Name,
		Description: Description,
		Status:      Status,
		Amount: Amount,
	}

	error := uc.plannerRepository.SaveEvent(newEvent, ctx)

	if error != nil {
		return error
	}

	return nil
}
