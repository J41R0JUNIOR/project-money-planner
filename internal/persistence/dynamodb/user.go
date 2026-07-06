package dynamodb

import (
	"github.com/J41R0JUNIOR/project-money-planner/internal/domain"
)

type UserRepository struct {
	client *Client
}

func NewUserRepository(client *Client) *UserRepository {
	return &UserRepository{client: client}
}

func (r *UserRepository) Create(user *domain.User) error {
	_ = user
	return nil
}
