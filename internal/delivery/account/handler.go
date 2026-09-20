package account

import (
	"context"
	"encoding/json"

	usecase "money-manager/internal/application/account"

	"github.com/aws/aws-lambda-go/events"

	"money-manager/internal/delivery/account/dto"
)

type AccountHandler struct {
	createAccountUseCase *usecase.CreateAccountUseCase
	getAccountsByUserIdUseCase   *usecase.GetAccountsByUserIdUseCase
}

func NewHandler(createAccountUseCase *usecase.CreateAccountUseCase, getAccountsByUserId *usecase.GetAccountsByUserIdUseCase) *AccountHandler {
	return &AccountHandler{
		createAccountUseCase: createAccountUseCase,
		getAccountsByUserIdUseCase: getAccountsByUserId,
	}
}

func (h *AccountHandler) CreateAccount(ctx context.Context, event events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	var requestDTO dto.CreateAccountRequestDTO

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

		err := h.createAccountUseCase.Execute(ctx, userId, requestDTO.Name, requestDTO.Type, requestDTO.Balance)
		if err != nil {
			return events.APIGatewayV2HTTPResponse{
				StatusCode: 500,
				Body:       `{"message":"Internal server error"}`,
			}, err
		}


	return events.APIGatewayV2HTTPResponse{
		StatusCode: 201,
		Body:       `{"message":"Account created successfully"}`,
	}, nil
}

func (h *AccountHandler) GetAccounts(ctx context.Context, event events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	userId := event.RequestContext.Authorizer.JWT.Claims["sub"]

	if userId == "" {
		return events.APIGatewayV2HTTPResponse{
			StatusCode: 401,
			Body:       `{"message":"Unauthorized"}`,
		}, nil
	}

	accounts, err := h.getAccountsByUserIdUseCase.Execute(ctx, userId)
	if err != nil {     
		return events.APIGatewayV2HTTPResponse{
			StatusCode: 500,
			Body:       `{"message":"Internal server error"}`,
		}, err
	}

	responseBody, err := json.Marshal(accounts)
	if err != nil {
		return events.APIGatewayV2HTTPResponse{
			StatusCode: 500,
			Body:       `{"message":"Internal server error"}`,
		}, err
	}

	return events.APIGatewayV2HTTPResponse{
		StatusCode: 200,
		Body:       string(responseBody),
	}, nil
}