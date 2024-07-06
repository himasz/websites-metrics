package metrics

import (
	"websites_metrics/config"
	"websites_metrics/models"
)

type IURLMetricsCalculator interface {
	CalculateMetrics(urlConfig models.URLConfig) (models.Metric, error)
	RunMetricsCalculations(cfg config.URLsConfig)
}
