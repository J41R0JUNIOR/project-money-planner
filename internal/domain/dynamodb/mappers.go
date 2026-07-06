package dynamodb_domain

import (
	"fmt"
	"time"

	base "github.com/J41R0JUNIOR/project-money-planner/internal/domain"
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

func NewUserItem(user *base.User) *UserItem {
	return &UserItem{
		BaseItem: BaseItem{
			PK:         BuildPK(user.ID),
			SK:         BuildUserSK(),
			EntityType: "USER",
		},
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}
}

func ToUserDomain(item *UserItem) *base.User {
	return &base.User{
		ID:       item.ID,
		Username: item.Username,
		Email:    item.Email,
	}
}

func NewTransactionItem(userID int, tx *base.Transaction) *TransactionItem {
	date, err := time.Parse("2006-01-02", tx.Date)
	if err != nil {
		date = time.Now().UTC()
	}

	return &TransactionItem{
		BaseItem: BaseItem{
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

func ToTransactionDomain(item *TransactionItem) *base.Transaction {
	return &base.Transaction{
		ID:          item.ID,
		Type:        item.Type,
		Amount:      item.Amount,
		Date:        item.Date,
		Description: item.Description,
		IsRecurring: item.IsRecurring,
	}
}

func NewProjectionItem(userID int, projection *base.Projection) *ProjectionItem {
	return &ProjectionItem{
		BaseItem: BaseItem{
			PK:         BuildPK(userID),
			SK:         BuildProjectionSK(projection.Date),
			EntityType: "PROJECTION",
		},
		Date:    projection.Date,
		Balance: projection.Balance,
	}
}

func ToProjectionDomain(item *ProjectionItem) *base.Projection {
	return &base.Projection{
		Date:    item.Date,
		Balance: item.Balance,
	}
}
