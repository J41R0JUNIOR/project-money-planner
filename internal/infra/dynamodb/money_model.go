package dynamodb

type MoneyItem struct {
	Amount   int64  `dynamodbav:"Amount"`
	Currency string `dynamodbav:"Currency"`
}