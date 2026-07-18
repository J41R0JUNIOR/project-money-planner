package dynamodb

import (
	"context"

	"money-manager/internal/domain/user"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type UserRepository struct {
	client    *dynamodb.Client
	tableName string
}

func NewUserRepository(client *dynamodb.Client, tableName string) *UserRepository {
	return &UserRepository{
		client:    client,
		tableName: tableName,
	}
}

func (r UserRepository) Save(user domain.User, ctx context.Context) error {
	item := UserItem{
		PK: "USER#" + user.Id,
		SK: "PROFILE",

		Id:    user.Id,
		Name:  user.Name,
		Email: user.Email,
	}

	av, err := attributevalue.MarshalMap(item)

	_, err = r.client.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(r.tableName),
		Item:      av,
	})

	return err
}
