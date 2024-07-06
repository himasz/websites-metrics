package repository

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
	"websites_metrics/models"
)

func TestSave(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect database: %v", err)
	}

	db.AutoMigrate(&models.Metric{})

	repo := NewMetricsRepository(db)

	metric := models.Metric{
		URL:          "http://example.com",
		Timestamp:    "2024-07-06T20:12:48+02:00",
		ResponseTime: 0.225,
		StatusCode:   200,
		RegexMatch:   true,
	}

	err = repo.Save(metric)
	if err != nil {
		t.Fatalf("Failed to save metric: %v", err)
	}

	var count int64
	db.Model(&models.Metric{}).Count(&count)
	if count != 1 {
		t.Errorf("Expected 1 metric, got %d", count)
	}
}
