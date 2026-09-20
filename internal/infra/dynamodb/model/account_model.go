package dynamodb_model

import domain "money-manager/internal/domain/account"

type Account struct {
	PK      string             `dynamodbav:"PK"`
	SK      string             `dynamodbav:"SK"`
	Id      string             `dynamodbav:"Id"`
	UserId  string             `dynamodbav:"UserId"`
	Name    string             `dynamodbav:"Name"`
	Type    domain.AccountType `dynamodbav:"Type"`
	Balance MoneyItem          `dynamodbav:"Balance"`
}
