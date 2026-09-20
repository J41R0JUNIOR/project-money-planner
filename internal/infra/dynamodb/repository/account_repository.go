package repository

import (
	"context"
	domain "money-manager/internal/domain/account"
	dynamodb_model "money-manager/internal/infra/dynamodb/model"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type AccountRepository struct {
	client    *dynamodb.Client
	tableName string
}

func NewAccountRepository(client *dynamodb.Client, tableName string) *AccountRepository {
	return &AccountRepository{
		client:    client,
		tableName: tableName,
	}
}
 
func (r *AccountRepository) SaveAccount(account domain.Account, ctx context.Context) error {
	newAccount := dynamodb_model.Account{
		PK:      "USER#" + account.UserId,
		SK:      "ACCOUNT#" + account.Id,
		Id:      account.Id,
		UserId:  account.UserId,
		Name:    account.Name,
		Type:    account.Type,
		Balance: dynamodb_model.MoneyItem{
			Amount: account.Balance.Amount, 
			Currency: string(account.Balance.Currency),
		},
	}

	av, err := attributevalue.MarshalMap(newAccount)
	if err != nil {
		return err
	}

	_, err = r.client.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: &r.tableName,
		Item:      av,
	})

	if err != nil {
		return err
	}

	return nil
}

func (r *AccountRepository) GetAccountsByUserId(userId string, ctx context.Context) ([]domain.Account, error) {
	panic("unimplemented")
}