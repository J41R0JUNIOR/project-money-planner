package dynamodb_domain

import (
	"time"

	base "github.com/J41R0JUNIOR/project-money-planner/internal/domain"
)

type BaseItem struct {
	PK         string `dynamodbav:"pk" json:"PK"`
	SK         string `dynamodbav:"sk" json:"SK"`
	EntityType string `dynamodbav:"entity_type" json:"EntityType"`
}

type UserItem struct {
	BaseItem
	ID       int    `dynamodbav:"id" json:"ID"`
	Username string `dynamodbav:"username" json:"Username"`
	Email    string `dynamodbav:"email" json:"Email"`
}

type TransactionItem struct {
	BaseItem
	ID          int                  `dynamodbav:"id" json:"ID"`
	UserID      int                  `dynamodbav:"user_id" json:"UserID"`
	Type        base.TypeTransaction `dynamodbav:"type" json:"Type"`
	Amount      float64              `dynamodbav:"amount" json:"Amount"`
	Date        string               `dynamodbav:"date" json:"Date"`
	Description string               `dynamodbav:"description" json:"Description"`
	IsRecurring bool                 `dynamodbav:"is_recurring" json:"IsRecurring"`
}

type ProjectionItem struct {
	BaseItem
	Date    time.Time `dynamodbav:"date" json:"Date"`
	Balance float64   `dynamodbav:"balance" json:"Balance"`
}
