package metrics

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
	"time"
	"websites_metrics/config"
	"websites_metrics/models"
	"websites_metrics/repository"
	"websites_metrics/scheduler"
)

func TestCalculateURLsMetrics(t *testing.T) {
	// Use SQLite in-memory database for integration test
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect database: %v", err)
	}

	db.AutoMigrate(&models.Metric{})

	repo := repository.NewMetricsRepository(db)
	cronScheduler := scheduler.NewCronScheduler()

	metricsCalculator := &URLMetricsCalculator{
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

	time.Sleep(2 * time.Second)

	var count int64
	db.Model(&models.Metric{}).Count(&count)
	if count == 0 {
		t.Errorf("Expected at least 1 metric, got %d", count)
	}
}
