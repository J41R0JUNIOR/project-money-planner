package planner

import (
	"context"
	"time"

	"github.com/google/uuid"

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
	StartDate string,
	Recurrence *domain.Recurrence,
) error {
	startDateParsed, err := time.Parse(time.RFC3339, StartDate)
	if err != nil {
		return err
	}

	var newEvent = domain.PlannedEvent{
		Id:          uuid.New().String(),
		UserId:      UserId,
		AccountId:   AccountId,
		CategoryId:  CategoryId,
		Name:        Name,
		Description: Description,
		Status:      Status,
		Amount:      Amount,
		StartDate:   startDateParsed,
		Recurrence:  Recurrence,
	}

	error := uc.plannerRepository.SaveEvent(newEvent, ctx)

	if error != nil {
		return error
	}

	return nil
}
