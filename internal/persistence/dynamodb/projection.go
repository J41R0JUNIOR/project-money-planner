package dynamodb

import (
	"github.com/J41R0JUNIOR/project-money-planner/internal/domain"
)

type ProjectionRepository struct {
	client *Client
}

func NewProjectionRepository(client *Client) *ProjectionRepository {
	return &ProjectionRepository{client: client}
}

func (r *ProjectionRepository) Create(projection *domain.Projection) error {
	_ = projection
	return nil
}
