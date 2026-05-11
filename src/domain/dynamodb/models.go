package dynamodb_model

import (
	"time"

	base "github.com/J41R0JUNIOR/project-money-planner/src/domain"
)

type BaseItem struct {
	PK         string `json:"PK"`
	SK         string `json:"SK"`
	EntityType string `json:"EntityType"`
}

type UserItem struct {
	BaseItem
	ID       int    `json:"ID"`
	Username string `json:"Username"`
	Email    string `json:"Email"`
	Password string `json:"Password"`
}

type TransactionItem struct {
	BaseItem
	ID          int                  `json:"ID"`
	UserID      int                  `json:"UserID"`
	Type        base.TypeTransaction `json:"Type"`
	Amount      float64              `json:"Amount"`
	Date        string               `json:"Date"`
	Description string               `json:"Description"`
	IsRecurring bool                 `json:"IsRecurring"`
}

type ProjectionItem struct {
	BaseItem
	Date    time.Time `json:"Date"`
	Balance float64   `json:"Balance"`
}
