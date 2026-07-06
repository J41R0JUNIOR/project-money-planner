package aws

import (
	"context"
	"fmt"

	awsinfra "github.com/J41R0JUNIOR/project-money-planner/internal/aws"
	awsconfig "github.com/aws/aws-sdk-go-v2/aws"
	awscfg "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
)

type ClientFactory struct {
	cfg awsconfig.Config
}

func NewClientFactory(ctx context.Context, cfg awsinfra.Config) (*ClientFactory, error) {
	loaded, err := awscfg.LoadDefaultConfig(ctx, awscfg.WithRegion(cfg.Region))
	if err != nil {
		return nil, fmt.Errorf("load aws config: %w", err)
	}

	return &ClientFactory{cfg: loaded}, nil
}

// NewDynamoDBClient cria um cliente DynamoDB.
func (f *ClientFactory) NewDynamoDBClient() *dynamodb.Client {
	return dynamodb.NewFromConfig(f.cfg)
}

// NewLambdaClient cria um cliente Lambda.
func (f *ClientFactory) NewLambdaClient() *lambda.Client {
	return lambda.NewFromConfig(f.cfg)
}
