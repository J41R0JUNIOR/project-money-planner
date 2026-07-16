package auth

import (
	"context"
	"encoding/json"
	"net/http"

	"money-manager/internal/api/auth/dto"
	"money-manager/internal/service"

	"github.com/aws/aws-lambda-go/events"
)

type Handler struct {
	authService *service.AuthService
}

func NewHandler(authService *service.AuthService) *Handler {
	return &Handler{
		authService: authService,
	}
}

func (h *Handler) Handle(
	ctx context.Context,
	event events.APIGatewayV2HTTPRequest,
) (events.APIGatewayV2HTTPResponse, error) {

	switch {
	case event.RequestContext.HTTP.Method == http.MethodPost &&
		event.RawPath == "/signup":

		return h.SignUp(ctx, event)

	case event.RequestContext.HTTP.Method == http.MethodPost &&
		event.RawPath == "/signin":

		return h.SignIn(ctx, event)

	default:

		return events.APIGatewayV2HTTPResponse{
			StatusCode: http.StatusNotFound,
			Body:       `{"message":"route not found"}`,
		}, nil
	}
}

func (h *Handler) SignUp(
	ctx context.Context,
	event events.APIGatewayV2HTTPRequest,
) (events.APIGatewayV2HTTPResponse, error) {

	requestDTO := auth_dto.SignUpRequestDTO{}

	if err := json.Unmarshal([]byte(event.Body), &requestDTO); err != nil {
		return events.APIGatewayV2HTTPResponse{
			StatusCode: http.StatusBadRequest,
			Body:       `{"message":"Invalid request body"}`,
		}, nil
	}

	err := h.authService.SignUp(ctx, requestDTO)

	if err != nil {
		return events.APIGatewayV2HTTPResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       err.Error(),
		}, nil
	}

	return events.APIGatewayV2HTTPResponse{
		StatusCode: http.StatusCreated,
		Body:       `{"message":"SignUp successful"}`,
	}, nil
}

func (h *Handler) SignIn(ctx context.Context, event events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {

	var requestDTO auth_dto.SignInRequestDTO

	if err := json.Unmarshal([]byte(event.Body), &requestDTO); err != nil {
		return events.APIGatewayV2HTTPResponse{
			StatusCode: http.StatusBadRequest,
			Body:       `{"message":"Invalid request body"}`,
		}, nil
	}

	token, err := h.authService.SignIn(ctx, requestDTO)

	if err != nil {
		return events.APIGatewayV2HTTPResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       err.Error(),
		}, nil
	}

	responseBody, _ := json.Marshal(map[string]string{
		"token": token,
	})

	return events.APIGatewayV2HTTPResponse{
		StatusCode: http.StatusOK,
		Body:       string(responseBody),
	}, nil
}