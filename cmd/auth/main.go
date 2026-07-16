package main

import (
	"context"

	"money-manager/internal/api/auth"
	"money-manager/internal/service"
	"money-manager/internal/infra/dynamodb"
	"money-manager/internal/infra/cognito"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/config"
)

func main() {
	awsConfig, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		panic(err)
	}

	cognitoClient := cognitoidentityprovider.NewFromConfig(awsConfig)

	authProvider := cognito.NewCognitoActions(cognitoClient)

	userRepository := dynamodb.NewUserRepository()

	service := service.NewAuthService(authProvider, userRepository)

	handler := auth.NewHandler(service)

	lambda.Start(handler.Handle)
}