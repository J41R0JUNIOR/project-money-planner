package user

import (
	"context"
	"encoding/json"
	"money-manager/internal/api/user/dto"
	"money-manager/internal/service"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

type Handler struct {
	userService *service.UserService
}

func NewHandler(userService *service.UserService) *Handler {
	return &Handler{
		userService: userService,
	}
}

func (h *Handler) Handle(
	ctx context.Context,
	event events.APIGatewayV2HTTPRequest,
) (events.APIGatewayV2HTTPResponse, error) {

	switch {
	case event.RequestContext.HTTP.Method == http.MethodPost &&
		event.RawPath == "/users":

		return h.CreateUser(ctx, event)

	default:

		return events.APIGatewayV2HTTPResponse{
			StatusCode: http.StatusNotFound,
			Body:       `{"message":"route not found"}`,
		}, nil
	}
}

func (h *Handler) CreateUser(ctx context.Context, event events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	var requestDTO dto.CreateUserRequestDTO

	if err := json.Unmarshal([]byte(event.Body), &requestDTO); err != nil {
		return events.APIGatewayV2HTTPResponse{
			StatusCode: http.StatusBadRequest,
			Body:       `{"message":"Invalid request body"}`,
		}, nil
	}
	return events.APIGatewayV2HTTPResponse{
		StatusCode: http.StatusCreated,
		Body:       `{"message":"User created successfully"}`,
	}, nil
}