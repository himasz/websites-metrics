package repository

import "websites_metrics/models"

type IMetricsRepository interface {
	Save(metrics models.Metric) error
}
