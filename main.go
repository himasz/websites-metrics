package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"websites_metrics/config"
	"websites_metrics/metrics"
	"websites_metrics/models"
	"websites_metrics/repository"
	"websites_metrics/scheduler"
)

func main() {
	var configLoader config.ILoader = &config.JSONLoader{}

	var dbCfg config.DBConfig
	err := configLoader.Load("config/json/db_config.json", &dbCfg)
	if err != nil {
		log.Fatalf("failed to load database config: %v", err)
	}

	var urlsCfg config.URLsConfig
	err = configLoader.Load("config/json/urls_config.json", &urlsCfg)
	if err != nil {
		log.Fatalf("failed to load URLs config: %v", err)
	}

	dsn := fmt.Sprintf("user=%s password=%s dbname=%s port=%d",
		dbCfg.Database.User, dbCfg.Database.Password, dbCfg.Database.DBName, dbCfg.Database.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	err = db.AutoMigrate(&models.Metric{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	var repo repository.IMetricsRepository = repository.NewMetricsRepository(db)
	var cronScheduler scheduler.IScheduler = scheduler.NewCronScheduler()

	metricsCalculator := &metrics.URLMetricsCalculator{
		Repo:       repo,
		IScheduler: cronScheduler,
	}

	metricsCalculator.RunMetricsCalculations(urlsCfg)

	cronScheduler.Start()

	select {}
}
