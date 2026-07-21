package delivery

import (
	"context"
	"encoding/json"
	usecase "money-manager/internal/application/planner"
	"money-manager/internal/delivery/planner/dto"

	"github.com/aws/aws-lambda-go/events"
)

type PlannerHandler struct {
	createEventUseCase *usecase.CreateEventUseCase
	deleteEventUseCase *usecase.DeleteEventUseCase
	updateEventUseCase *usecase.UpdateEventUseCase
}

func (h *PlannerHandler) NewHandler(
	createEventUseCase *usecase.CreateEventUseCase,
	deleteEventUseCase *usecase.DeleteEventUseCase,
	updateEventUseCase *usecase.UpdateEventUseCase,

) *PlannerHandler {

	return &PlannerHandler{
		createEventUseCase: createEventUseCase,
		deleteEventUseCase: deleteEventUseCase,
		updateEventUseCase: updateEventUseCase,
	}
}

func (h *PlannerHandler) CreateEvent(
	ctx context.Context,
	event events.APIGatewayV2HTTPRequest,
) (events.APIGatewayV2HTTPResponse, error) {

	var requestDTO dto.CreateEventRequestDTO

	if err := json.Unmarshal([]byte(event.Body), &requestDTO); err != nil {
		return events.APIGatewayV2HTTPResponse{
			StatusCode: 400,
			Body:       `{"message":"Invalid request body"}`,
		}, err
	}

	error := h.createEventUseCase.Execute(ctx, requestDTO.UserId, requestDTO.AccountId, requestDTO.CategoryId, requestDTO.Name, requestDTO.Description, requestDTO.PlannedEventStatus, requestDTO.Amount)

	if error != nil {
		return events.APIGatewayV2HTTPResponse{
			StatusCode: 500,
			Body:       `{"message":"` + error.Error() + `"}`,
		}, nil
	}

	return events.APIGatewayV2HTTPResponse{
		StatusCode: 200,
		Body:       `{"message":"Event created successfully"}`,
	}, nil
}

func (h *PlannerHandler) DeleteEvent(
	ctx context.Context,
	event events.APIGatewayV2HTTPRequest,
) (events.APIGatewayV2HTTPResponse, error) {

	if event.PathParameters == nil || event.PathParameters["event_id"] == "" {
		return events.APIGatewayV2HTTPResponse{
			StatusCode: 400,
			Body:       `{"message":"Invalid request body"}`,
		}, nil
	}

	error := h.deleteEventUseCase.Execute(ctx, event.PathParameters["event_id"])

	if error != nil {
		return events.APIGatewayV2HTTPResponse{
			StatusCode: 500,
			Body:       `{"message":"` + error.Error() + `"}`,
		}, nil
	}

	return events.APIGatewayV2HTTPResponse{
		StatusCode: 200,
		Body:       `{"message":"Event deleted successfully"}`,
	}, nil
}

func (h *PlannerHandler) UpdateEvent(
	ctx context.Context,
	event events.APIGatewayV2HTTPRequest,
) (events.APIGatewayV2HTTPResponse, error) {

	var requestDTO dto.UpdateEventRequestDTO
	
	if err := json.Unmarshal([]byte(event.Body), &requestDTO); err != nil {
		return events.APIGatewayV2HTTPResponse{
			StatusCode: 400,
			Body:       `{"message":"Invalid request body"}`,
		}, err
	}

	error := h.updateEventUseCase.Execute(ctx, requestDTO.ToDomain())

	if error != nil {
		return events.APIGatewayV2HTTPResponse{
			StatusCode: 500,
			Body:       `{"message":"` + error.Error() + `"}`,
		}, nil
	}

	return events.APIGatewayV2HTTPResponse{
		StatusCode: 200,
		Body:       `{"message":"Event updated successfully"}`,
	}, nil
}