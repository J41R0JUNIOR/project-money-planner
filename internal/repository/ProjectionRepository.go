package repository

import "github.com/J41R0JUNIOR/project-money-planner/internal/domain"

type ProjectionRepository interface {
	CreateProjection(projection *domain.Projection) error
	GetProjectionByDate(date string) (*domain.Projection, error)
	ListProjections() ([]*domain.Projection, error)
	DeleteProjection(date string) error
}
