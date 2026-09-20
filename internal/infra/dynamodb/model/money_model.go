package dynamodb_model

import "money-manager/internal/domain"

type MoneyItem struct {
	Amount   int64  `dynamodbav:"Amount"`
	Currency string `dynamodbav:"Currency"`
}

func (m MoneyItem) toDomain() domain.Money {
	return domain.Money{
		Amount:   m.Amount,
		Currency: domain.Currency(m.Currency),
	}
}