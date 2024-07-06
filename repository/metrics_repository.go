package repository

import (
	"gorm.io/gorm"
	"websites_metrics/models"
)

type MetricsRepositoryImpl struct {
	db *gorm.DB
}

func NewMetricsRepositoryImpl(db *gorm.DB) *MetricsRepositoryImpl {
	return &MetricsRepositoryImpl{db: db}
}

func (r *MetricsRepositoryImpl) Save(metrics models.Metric) error {
	return r.db.Create(&metrics).Error
}
