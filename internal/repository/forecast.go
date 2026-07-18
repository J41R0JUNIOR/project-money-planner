package repository

import (
	domain "money-manager/internal/domain/forecast"
	"context"
)

type ForecastRepository interface {
	SaveScenario(scenario domain.Scenario, ctx context.Context) error
	DeleteScenario(scenarioID string, ctx context.Context) error
}