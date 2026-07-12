package auth

import (
	"context"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	dto "github.com/J41R0JUNIOR/project-money-planner/internal/lambdas/auth/dto"
)

func SignUpHandler(event events.APIGatewayV2HTTPRequest, ctx context.Context) (events.APIGatewayV2HTTPResponse, error) {
	dto := SignUpRequestDTO{}

	return events.APIGatewayV2HTTPResponse{
		StatusCode: http.StatusOK,
		Body:	   `{"message": "SignUp successful"}`,
	}, nil
}


func SignInHandler(event events.APIGatewayV2HTTPRequest, ctx context.Context) (events.APIGatewayV2HTTPResponse, error) {
	return events.APIGatewayV2HTTPResponse{
		StatusCode: http.StatusOK,
		Body:	   `{"message": "SignIn successful"}`,
	}, nil
}