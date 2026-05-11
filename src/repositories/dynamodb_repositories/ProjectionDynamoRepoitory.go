package dynamodb_repositories

import (
	"github.com/J41R0JUNIOR/project-money-planner/src/domain"
	"github.com/J41R0JUNIOR/project-money-planner/src/repositories"
)

type ProjectionDynamoRepository struct {
	// DynamoDB client and other necessary fields
}

// CreateProjection implements [repositories.ProjectionRepository].
func (p *ProjectionDynamoRepository) CreateProjection(projection *domain.Projection) error {
	panic("unimplemented")
}

// DeleteProjection implements [repositories.ProjectionRepository].
func (p *ProjectionDynamoRepository) DeleteProjection(date string) error {
	panic("unimplemented")
}

// GetProjectionByDate implements [repositories.ProjectionRepository].
func (p *ProjectionDynamoRepository) GetProjectionByDate(date string) (*domain.Projection, error) {
	panic("unimplemented")
}

// ListProjections implements [repositories.ProjectionRepository].
func (p *ProjectionDynamoRepository) ListProjections() ([]*domain.Projection, error) {
	panic("unimplemented")
}

// Ensure ProjectionDynamoRepository implements the ProjectionRepository interface
var _ repositories.ProjectionRepository = (*ProjectionDynamoRepository)(nil)
