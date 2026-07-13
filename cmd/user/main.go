package main

import (
	"money-manager/internal/api/user"
	"money-manager/internal/service"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {

	userService := service.NewUserService()

	handler := user.NewHandler(userService)

	lambda.Start(handler.Handle)
}