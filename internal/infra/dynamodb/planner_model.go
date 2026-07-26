package dynamodb

type EventItem struct {
	PK                 string `dynamodbav:"PK"`
	SK                 string `dynamodbav:"SK"`
	Id                 string `dynamodbav:"Id"`
	UserId             string `dynamodbav:"UserId"`
	AccountId          string `dynamodbav:"AccountId"`
	CategoryId         string `dynamodbav:"CategoryId"`
	Name               string `dynamodbav:"Name"`
	Description        string `dynamodbav:"Description"`
	PlannedEventStatus string `dynamodbav:"PlannedEventStatus"`
	Amount             MoneyItem  `dynamodbav:"Amount"`
}