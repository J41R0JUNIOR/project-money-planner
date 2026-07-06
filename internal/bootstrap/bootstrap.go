package bootstrap

import (
	"context"
	"fmt"

	awsinfra "github.com/J41R0JUNIOR/project-money-planner/internal/aws"
	persistencedynamodb "github.com/J41R0JUNIOR/project-money-planner/internal/persistence/dynamodb"
)

type Application struct {
	AWSFactory *awsinfra.ServiceFactory
	DynamoDB   *persistencedynamodb.Client
}

func New(ctx context.Context) (*Application, error) {
	cfg := awsinfra.LoadConfig()
	factory, err := awsinfra.NewServiceFactory(ctx, cfg)
	if err != nil {
		return nil, fmt.Errorf("create aws service factory: %w", err)
	}

	db, err := persistencedynamodb.NewClient(ctx, factory)
	if err != nil {
		return nil, fmt.Errorf("create dynamodb client: %w", err)
	}

	return &Application{AWSFactory: factory, DynamoDB: db}, nil
}
