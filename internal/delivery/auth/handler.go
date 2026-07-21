package auth

import (
	"context"
	"encoding/json"
	usecase "money-manager/internal/application/auth"
	"money-manager/internal/delivery/auth/dto"

	"github.com/aws/aws-lambda-go/events"
)

type Handler struct {
	signUpUseCase      *usecase.SignUpUseCase
	signInUseCase      *usecase.SignInUseCase
	confirmCodeUseCase *usecase.ConfirmCodeUseCase
	refreshUseCase     *usecase.RefreshUseCase
}

func NewHandler(
	signUp *usecase.SignUpUseCase,
	confirmCode *usecase.ConfirmCodeUseCase,
	signIn *usecase.SignInUseCase,
	refresh *usecase.RefreshUseCase,

) *Handler {

	return &Handler{
		signUpUseCase:      signUp,
		signInUseCase:      signIn,
		confirmCodeUseCase: confirmCode,
		refreshUseCase:     refresh,
	}
}

func (h *Handler) SignUp(
	ctx context.Context,
	event events.APIGatewayV2HTTPRequest,
) (events.APIGatewayV2HTTPResponse, error) {

	var requestDTO dto.SignUpRequestDTO

	if err := json.Unmarshal([]byte(event.Body), &requestDTO); err != nil {
		return events.APIGatewayV2HTTPResponse{
			StatusCode: 400,
			Body:       `{"message":"Invalid request body"}`,
		}, err
	}

	error := h.signUpUseCase.Execute(ctx, requestDTO.Name, requestDTO.Email, requestDTO.Password)

	if error != nil {
		return events.APIGatewayV2HTTPResponse{
			StatusCode: 500,
			Body:       `{"message":"` + error.Error() + `"}`,
		}, nil
	}

	return events.APIGatewayV2HTTPResponse{
		StatusCode: 200,
		Body:       "User signed up successfully",
	}, nil

}

func (h *Handler) SignIn(
	ctx context.Context,
	event events.APIGatewayV2HTTPRequest,
) (events.APIGatewayV2HTTPResponse, error) {

	var requestDTO dto.SignInRequestDTO
	if err := json.Unmarshal([]byte(event.Body), &requestDTO); err != nil {
		return events.APIGatewayV2HTTPResponse{
			StatusCode: 400,
			Body:       `{"message":"Invalid request body"}`,
		}, err
	}

	response, err := h.signInUseCase.Execute(ctx, requestDTO.Email, requestDTO.Password)

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

func (h *Handler) ConfirmCode(
	ctx context.Context,
	event events.APIGatewayV2HTTPRequest,
) (events.APIGatewayV2HTTPResponse, error) {
	var requestDTO dto.ConfirmCodeRequestDTO
	if err := json.Unmarshal([]byte(event.Body), &requestDTO); err != nil {
		return events.APIGatewayV2HTTPResponse{
			StatusCode: 400,
			Body:       `{"message":"Invalid request body"}`,
		}, err
	}

	response, err := h.confirmCodeUseCase.Execute(ctx, requestDTO.Email, requestDTO.Code)

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

func (h *Handler) RefreshToken(
	ctx context.Context,
	event events.APIGatewayV2HTTPRequest,
) (events.APIGatewayV2HTTPResponse, error) {
	var requestDTO dto.RefreshRequestDTO
	if err := json.Unmarshal([]byte(event.Body), &requestDTO); err != nil {
		return events.APIGatewayV2HTTPResponse{
			StatusCode: 400,
			Body:       `{"message":"Invalid request body"}`,
		}, err
	}

	response, err := h.refreshUseCase.Execute(ctx, requestDTO.RefreshToken)

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
