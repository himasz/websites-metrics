package repository

import (
	"gorm.io/gorm"
	"websites_metrics/models"
)

type MetricsRepository struct {
	db *gorm.DB
}

func NewMetricsRepository(db *gorm.DB) *MetricsRepository {
	return &MetricsRepository{db: db}
}

func (r *MetricsRepository) Save(metrics models.Metric) error {
	return r.db.Create(&metrics).Error
}

var _ IMetricsRepository = (*MetricsRepository)(nil)
