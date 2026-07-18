package planner

import (
	"context"

	domain "money-manager/internal/domain/planner"
	root "money-manager/internal/domain"
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
    AccountId  string,
    CategoryId string,
    Name        string,
    Description string,
    Status domain.PlannedEventStatus,
    Amount root.Money,
) error {
	var newEvent = domain.PlannedEvent{
		Id: "",
		UserId: UserId,
		AccountId: AccountId,
		CategoryId: CategoryId,
		Name: Name,
		Description: Description,
		Status:	Status,
		Amount: root.Money{
			Amount: 0,
			Currency: "BRL",
		},
	}

	error := uc.plannerRepository.SaveEvent(newEvent, ctx)

	if error != nil {
		return error
	}

	return nil
}