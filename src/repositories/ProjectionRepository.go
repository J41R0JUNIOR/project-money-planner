package repositories

import "github.com/J41R0JUNIOR/project-money-planner/src/domain"

type ProjectionRepository interface {
	createProjection(projection *domain.Projection) error
	getProjectionByDate(date string) (*domain.Projection, error)
	listProjections() ([]*domain.Projection, error)
	deleteProjection(date string) error
}
