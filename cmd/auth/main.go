package main

import (
	"money-manager/internal/api/auth"
	"money-manager/internal/service"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {

	authService := service.NewAuthService()

	handler := auth.NewHandler(authService)

	lambda.Start(handler.Handle)
}