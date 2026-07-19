package main

import (
	"context"
	"os"

	usecase "money-manager/internal/application/auth"
	delivery "money-manager/internal/delivery/auth"
	provider "money-manager/internal/infra/cognito"
	db "money-manager/internal/infra/dynamodb"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func main() {

	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		panic(err)
	}

	clientID := os.Getenv("COGNITO_CLIENT_ID")
    if clientID == "" {
        panic("COGNITO_CLIENT_ID is not set")
    }

	tableName := os.Getenv("DYNAMODB_TABLE_NAME")
    if tableName == "" {
        panic("DYNAMODB_TABLE_NAME is not set")
    }

	cognitoClient := cognitoidentityprovider.NewFromConfig(cfg)
	dynamoClient := dynamodb.NewFromConfig(cfg)

	authProvider := provider.NewCognitoActions(cognitoClient, clientID)

	userRepository := db.NewUserRepository(dynamoClient, tableName)

	signUpUseCase := usecase.NewSignUpUseCase(
		authProvider,
		userRepository,
	)

	signInUseCase := usecase.NewSignInUseCase(
		authProvider,
	)

	refreshUseCase := usecase.NewRefreshUseCase(
		authProvider,
	)

	confirmCodeUseCase := usecase.NewConfirmCodeUseCase(
		authProvider,
	)

	handler := delivery.NewHandler(
		signUpUseCase,
		confirmCodeUseCase,
		signInUseCase,
		refreshUseCase,
	)

	router := delivery.NewRouter(handler)

	// Lambda
	lambda.Start(router.Handle)
}