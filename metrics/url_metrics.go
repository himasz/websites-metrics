package metrics

import (
	"io"
	"log"
	"net/http"
	"regexp"
	"time"
	"websites_metrics/config"
	"websites_metrics/models"
	"websites_metrics/repository"
)

type URLCheckerImpl struct {
	Repo *repository.MetricsRepositoryImpl
}

func (u *URLCheckerImpl) CalculateURLsMetrics(cfg config.URLsConfig) {
	for _, urlConfig := range cfg.URLs {
		go func(urlConfig config.URLConfig) {
			metrics, err := u.Check(models.URLConfig{
				URL:   urlConfig.URL,
				Regex: urlConfig.Regex,
			})
			if err != nil {
				log.Printf("error checking URL %s: %v", urlConfig.URL, err)
				return
			}
			log.Printf("URL Metrics: %v", metrics)
			err = u.Repo.Save(metrics)
			if err != nil {
				log.Printf("error saving metrics for URL %s: %v", urlConfig.URL, err)
			}
		}(urlConfig)
	}
}

func (u *URLCheckerImpl) Check(urlConfig models.URLConfig) (models.Metric, error) {
	start := time.Now()
	resp, err := http.Get(urlConfig.URL)
	if err != nil {
		return models.Metric{}, err
	}
	defer resp.Body.Close()

	responseTime := time.Since(start).Seconds()
	regexMatch := false

	if urlConfig.Regex != "" {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return models.Metric{}, err
		}
		match, err := regexp.Match(urlConfig.Regex, bodyBytes)
		if err != nil {
			return models.Metric{}, err
		}
		regexMatch = match
	}

	return models.Metric{
		URL:          urlConfig.URL,
		Timestamp:    time.Now().Format(time.RFC3339),
		ResponseTime: responseTime,
		StatusCode:   resp.StatusCode,
		RegexMatch:   regexMatch,
	}, nil
}
