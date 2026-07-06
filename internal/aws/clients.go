package aws

import (
	"context"
	"fmt"

	awsconfig "github.com/aws/aws-sdk-go-v2/aws"
	awscfg "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

type ServiceFactory struct {
	cfg awsconfig.Config
}

func NewServiceFactory(ctx context.Context, cfg Config) (*ServiceFactory, error) {
	options := []func(*awscfg.LoadOptions) error{}

	if cfg.Region != "" {
		options = append(options, awscfg.WithRegion(cfg.Region))
	}
	if cfg.Profile != "" {
		options = append(options, awscfg.WithSharedConfigProfile(cfg.Profile))
	}

	loaded, err := awscfg.LoadDefaultConfig(ctx, options...)
	if err != nil {
		return nil, fmt.Errorf("load aws config: %w", err)
	}

	return &ServiceFactory{cfg: loaded}, nil
}

func (f *ServiceFactory) NewDynamoDBClient() *dynamodb.Client {
	return dynamodb.NewFromConfig(f.cfg)
}

func (f *ServiceFactory) NewLambdaClient() *lambda.Client {
	return lambda.NewFromConfig(f.cfg)
}

func (f *ServiceFactory) NewEventBridgeClient() *eventbridge.Client {
	return eventbridge.NewFromConfig(f.cfg)
}

func (f *ServiceFactory) NewSecretsManagerClient() *secretsmanager.Client {
	return secretsmanager.NewFromConfig(f.cfg)
}
