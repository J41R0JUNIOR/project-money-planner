package dynamodb_repositories

import (
	"fmt"
	"time"

	base "github.com/J41R0JUNIOR/project-money-planner/src/domain"
	model "github.com/J41R0JUNIOR/project-money-planner/src/domain/dynamodb"
)

func BuildPK(userID int) string {
	return fmt.Sprintf("USER#%d", userID)
}

func BuildTransactionSK(date time.Time, transactionID int) string {
	return fmt.Sprintf("DATE#%s#TX#%d", date.UTC().Format(time.RFC3339Nano), transactionID)
}

func BuildProjectionSK(date time.Time) string {
	return fmt.Sprintf("DATE#%s#PROJECTION", date.UTC().Format("2006-01-02"))
}

func BuildUserSK() string {
	return "PROFILE"
}

func NewUserItem(user *base.User) *model.UserItem {
	return &model.UserItem{
		BaseItem: model.BaseItem{
			PK:         BuildPK(user.ID),
			SK:         BuildUserSK(),
			EntityType: "USER",
		},
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
	}
}

func ToUserDomain(item *model.UserItem) *base.User {
	return &base.User{
		ID:       item.ID,
		Username: item.Username,
		Email:    item.Email,
		Password: item.Password,
	}
}

func NewTransactionItem(userID int, tx *base.Transaction) *model.TransactionItem {
	date, err := time.Parse("2006-01-02", tx.Date)
	if err != nil {
		date = time.Now().UTC()
	}

	return &model.TransactionItem{
		BaseItem: model.BaseItem{
			PK:         BuildPK(userID),
			SK:         BuildTransactionSK(date, tx.ID),
			EntityType: "TRANSACTION",
		},
		ID:          tx.ID,
		UserID:      userID,
		Type:        tx.Type,
		Amount:      tx.Amount,
		Date:        tx.Date,
		Description: tx.Description,
		IsRecurring: tx.IsRecurring,
	}
}

func ToTransactionDomain(item *model.TransactionItem) *base.Transaction {
	return &base.Transaction{
		ID:          item.ID,
		Type:        item.Type,
		Amount:      item.Amount,
		Date:        item.Date,
		Description: item.Description,
		IsRecurring: item.IsRecurring,
	}
}

func NewProjectionItem(userID int, projection *base.Projection) *model.ProjectionItem {
	return &model.ProjectionItem{
		BaseItem: model.BaseItem{
			PK:         BuildPK(userID),
			SK:         BuildProjectionSK(projection.Date),
			EntityType: "PROJECTION",
		},
		Date:    projection.Date,
		Balance: projection.Balance,
	}
}

func ToProjectionDomain(item *model.ProjectionItem) *base.Projection {
	return &base.Projection{
		Date:    item.Date,
		Balance: item.Balance,
	}
}
