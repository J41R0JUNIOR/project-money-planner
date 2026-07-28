package main

import (
	"context"
	usecase "money-manager/internal/application/planner"
	"os"

	delivery "money-manager/internal/delivery/planner"
	db "money-manager/internal/infra/dynamodb"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
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

	dynamodbClient := dynamodb.NewFromConfig(cfg)

	plannerRepository := db.NewPlannerRepository(dynamodbClient, tableName)

	createEventUseCase := usecase.NewCreateEventUseCase(plannerRepository)
	deleteEventUseCase := usecase.NewDeleteEventUseCase(plannerRepository)
	readEventUseCase := usecase.NewReadEventUseCase(plannerRepository)

	handler := delivery.NewHandler(createEventUseCase, deleteEventUseCase, readEventUseCase)

	router := delivery.NewRouter(handler)

	lambda.Start(router.Handle)
}
