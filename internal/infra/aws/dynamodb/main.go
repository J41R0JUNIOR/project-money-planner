package dynamodb

import (
	"context"
	"fmt"

	awsinfra "github.com/J41R0JUNIOR/project-money-planner/internal/aws"
	awsdynamodb "github.com/aws/aws-sdk-go-v2/service/dynamodb"
	awslambda "github.com/aws/aws-sdk-go-v2/service/lambda"
)

type Clients struct {
	DynamoDB *awsdynamodb.Client
	Lambda   *awslambda.Client
}

func NewClients(ctx context.Context, cfg awsinfra.Config) (*Clients, error) {
	factory, err := awsinfra.NewServiceFactory(ctx, cfg)
	if err != nil {
		return nil, fmt.Errorf("create aws service factory: %w", err)
	}

	return &Clients{
		DynamoDB: factory.NewDynamoDBClient(),
		Lambda:   factory.NewLambdaClient(),
	}, nil
}
