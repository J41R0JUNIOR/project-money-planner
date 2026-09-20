package dto

import (
	root "money-manager/internal/domain"
	domain "money-manager/internal/domain/account"
)

type CreateAccountRequestDTO struct {
	Id      string             `json:"id"`
	UserId  string             `json:"user_id"`
	Name    string             `json:"name"`
	Type    domain.AccountType `json:"type"`
	Balance root.Money         `json:"balance"`
}
