package dynamodb_repositories

import (
	"github.com/J41R0JUNIOR/project-money-planner/src/domain"
	"github.com/J41R0JUNIOR/project-money-planner/src/repositories"
)

type UserDynamoRepository struct {
	// DynamoDB client and other necessary fields
}

// GetUserByID implements [repositories.UserRepository].
func (u *UserDynamoRepository) GetUserByID(id int) (*domain.User, error) {
	panic("unimplemented")
}

// CreateUser implements [repositories.UserRepository].
func (u *UserDynamoRepository) CreateUser(user *domain.User) error {
	panic("unimplemented")
}

// DeleteUser implements [repositories.UserRepository].
func (u *UserDynamoRepository) DeleteUser(id int) error {
	panic("unimplemented")
}

// Ensure UserDynamoRepository implements the UserRepository interface
var _ repositories.UserRepository = (*UserDynamoRepository)(nil)
