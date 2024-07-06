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
	"websites_metrics/scheduler"
)

type URLMetricsCalculator struct {
	Repo       repository.IMetricsRepository
	IScheduler scheduler.IScheduler
}

func (u *URLMetricsCalculator) RunMetricsCalculations(cfg config.URLsConfig) {
	for _, urlConfig := range cfg.URLs {
		urlConfig := urlConfig
		_, err := u.IScheduler.AddFunc(urlConfig.Interval, func() {
			metrics, err := u.CalculateMetrics(models.URLConfig{
				URL:   urlConfig.URL,
				Regex: urlConfig.Regex,
			})
			if err != nil {
				log.Printf("error checking URL %s: %v", urlConfig.URL, err)
				return
			}
			log.Printf("URL Metrics: %s %s %f %d %t", metrics.URL, metrics.Timestamp, metrics.ResponseTime, metrics.StatusCode, metrics.RegexMatch)
			err = u.Repo.Save(metrics)
			if err != nil {
				log.Printf("error saving metrics for URL %s: %v", urlConfig.URL, err)
			}
		})
		if err != nil {
			log.Printf("error scheduling URL check for %s: %v", urlConfig.URL, err)
		}
	}
}

func (u *URLMetricsCalculator) CalculateMetrics(urlConfig models.URLConfig) (models.Metric, error) {
	start := time.Now()
	resp, err := http.Get(urlConfig.URL)
	if err != nil {
		return models.Metric{}, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatalf("failed to close connection: %v", err)
		}
	}(resp.Body)

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

var _ IMetricsCalculator = (*URLMetricsCalculator)(nil)
