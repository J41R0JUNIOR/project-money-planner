package projection

import "github.com/J41R0JUNIOR/project-money-planner/internal/domain"

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Create(projection *domain.Projection) error {
	_ = projection
	return nil
}

func (s *Service) GetByID(id int) (*domain.Projection, error) {
	_ = id
	return &domain.Projection{}, nil
}
