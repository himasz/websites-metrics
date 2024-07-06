package test

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
	"time"
	"websites_metrics/config"
	"websites_metrics/metrics"
	"websites_metrics/models"
	"websites_metrics/repository"
	"websites_metrics/scheduler"
)

func TestMetricsCalculationIT(t *testing.T) {
	// Use SQLite in-memory database for integration test
	dsn := "file::memory:?cache=shared"
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}

	err = db.AutoMigrate(&models.Metric{})
	if err != nil {
		t.Fatalf("Failed to migrate database: %v", err)
	}

	repo := repository.NewMetricsRepository(db)
	cronScheduler := scheduler.NewCronScheduler()

	metricsCalculator := &metrics.URLMetricsCalculator{
		Repo:       repo,
		IScheduler: cronScheduler,
	}

	urlsConfig := config.URLsConfig{
		URLs: []config.URLConfig{
			{
				URL:      "http://example.com",
				Regex:    "Example Domain",
				Interval: "@every 1s",
			},
		},
	}

	metricsCalculator.RunMetricsCalculations(urlsConfig)
	cronScheduler.Start()

	// Wait some time for scheduled tasks to execute
	time.Sleep(5 * time.Second)

	var count int64
	db.Model(&models.Metric{}).Count(&count)
	if count == 0 {
		t.Errorf("Expected at least 1 metric, got %d", count)
	}
}
