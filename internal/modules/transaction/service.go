package transaction

import "github.com/J41R0JUNIOR/project-money-planner/internal/domain"

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Create(transaction *domain.Transaction) error {
	_ = transaction
	return nil
}

func (s *Service) GetByID(id int) (*domain.Transaction, error) {
	_ = id
	return &domain.Transaction{}, nil
}
