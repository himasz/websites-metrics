package models

import (
	"fmt"
	"gorm.io/gorm"
)

type Metric struct {
	gorm.Model
	URL          string
	Timestamp    string
	ResponseTime float64
	StatusCode   int
	RegexMatch   bool
}

type URLConfig struct {
	URL   string
	Regex string
}

func (m Metric) String() string {
	return fmt.Sprintf("%s %s %f %d %t", m.URL, m.Timestamp, m.ResponseTime, m.StatusCode, m.RegexMatch)
}
