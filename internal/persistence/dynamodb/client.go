package dynamodb

import (
	"context"
	"fmt"

	awsfactory "github.com/J41R0JUNIOR/project-money-planner/internal/aws"
	awsdynamodb "github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type Client struct {
	inner *awsdynamodb.Client
}

func NewClient(ctx context.Context, factory *awsfactory.ServiceFactory) (*Client, error) {
	if factory == nil {
		return nil, fmt.Errorf("aws service factory is nil")
	}

	return &Client{inner: factory.NewDynamoDBClient()}, nil
}

func (c *Client) Inner() *awsdynamodb.Client {
	return c.inner
}
