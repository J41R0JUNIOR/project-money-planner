package dynamodb

import (
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"money-manager/internal/domain"
)

type UserRepository struct {
	client *dynamodb.Client
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) Save(user domain.User) error {
	return nil
}

func (r *UserRepository) Delete(id string) error {
	panic("unimplemented")
}