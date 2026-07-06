package dynamodb_repositories

import (
	"github.com/J41R0JUNIOR/project-money-planner/internal/domain"
	"github.com/J41R0JUNIOR/project-money-planner/internal/repository"
)

type ProjectionDynamoRepository struct {
}

func (p *ProjectionDynamoRepository) CreateProjection(projection *domain.Projection) error {
	panic("unimplemented")
}

func (p *ProjectionDynamoRepository) DeleteProjection(date string) error {
	panic("unimplemented")
}

func (p *ProjectionDynamoRepository) GetProjectionByDate(date string) (*domain.Projection, error) {
	panic("unimplemented")
}

func (p *ProjectionDynamoRepository) ListProjections() ([]*domain.Projection, error) {
	panic("unimplemented")
}

var _ repository.ProjectionRepository = (*ProjectionDynamoRepository)(nil)
