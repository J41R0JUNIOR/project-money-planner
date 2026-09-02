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
	readEventUseCase   *usecase.ReadEventUseCase
}

func NewHandler(
	createEventUseCase *usecase.CreateEventUseCase,
	deleteEventUseCase *usecase.DeleteEventUseCase,
	readEventUseCase *usecase.ReadEventUseCase,
) *PlannerHandler {

	return &PlannerHandler{
		createEventUseCase: createEventUseCase,
		deleteEventUseCase: deleteEventUseCase,
		readEventUseCase:   readEventUseCase,
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

	userId := event.RequestContext.Authorizer.JWT.Claims["sub"]

	if userId == "" {
		return events.APIGatewayV2HTTPResponse{
			StatusCode: 401,
			Body:       `{"message":"Unauthorized"}`,
		}, nil
	}

	error := h.createEventUseCase.Execute(ctx, userId, requestDTO.AccountId, requestDTO.CategoryId, requestDTO.Name, requestDTO.Description, requestDTO.Status, requestDTO.Amount, requestDTO.StartDate, requestDTO.Recurrence)

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

	userId := event.RequestContext.Authorizer.JWT.Claims["sub"]

	if userId == "" {
		return events.APIGatewayV2HTTPResponse{
			StatusCode: 401,
			Body:       `{"message":"Unauthorized"}`,
		}, nil
	}

	error := h.deleteEventUseCase.Execute(ctx, userId, event.PathParameters["event_id"])

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

func (h *PlannerHandler) ReadEvent(
	ctx context.Context,
	event events.APIGatewayV2HTTPRequest,
) (events.APIGatewayV2HTTPResponse, error) {

	userId := event.RequestContext.Authorizer.JWT.Claims["sub"]

	if userId == "" {
		return events.APIGatewayV2HTTPResponse{
			StatusCode: 401,
			Body:       `{"message":"Unauthorized"}`,
		}, nil
	}

	response, err := h.readEventUseCase.Execute(ctx, userId)

	if err != nil {
		return events.APIGatewayV2HTTPResponse{
			StatusCode: 500,
			Body:       `{"message":"` + err.Error() + `"}`,
		}, nil
	}

	responseBody, _ := json.Marshal(response)

	return events.APIGatewayV2HTTPResponse{
		StatusCode: 200,
		Body:       string(responseBody),
	}, nil
}