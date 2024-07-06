package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"websites_metrics/config"
	"websites_metrics/metrics"
	"websites_metrics/models"
	"websites_metrics/repository"
	"websites_metrics/utils"
)

func main() {
	jsonLoader := &utils.JSONLoader{}

	var dbCfg config.DBConfig
	err := jsonLoader.LoadJson("db_config.json", &dbCfg)
	if err != nil {
		log.Fatalf("failed to load database config: %v", err)
	}

	var urlsCfg config.URLsConfig
	err = jsonLoader.LoadJson("urls.json", &urlsCfg)
	if err != nil {
		log.Fatalf("failed to load URLs config: %v", err)
	}

	// Construct the DSN
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s port=%d",
		dbCfg.Database.User, dbCfg.Database.Password, dbCfg.Database.DBName, dbCfg.Database.Port)

	// Connect to the database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// Perform migration
	err = db.AutoMigrate(&models.Metric{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	// Initialize repository
	repo := repository.NewMetricsRepositoryImpl(db)

	// Initialize URL checker
	metrics := &metrics.URLCheckerImpl{Repo: repo}

	// Initialize cron scheduler
	c := cron.New()
	c.AddFunc("@every 3s", func() {
		metrics.CalculateURLsMetrics(urlsCfg)
	})

	// Start cron scheduler
	c.Start()

	// Keep the main function running
	select {}
}
