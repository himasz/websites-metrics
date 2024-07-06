package scheduler

import (
	"github.com/robfig/cron/v3"
)

type CronScheduler struct {
	Cron *cron.Cron
}

func NewCronScheduler() *CronScheduler {
	return &CronScheduler{Cron: cron.New()}
}

func (s *CronScheduler) AddFunc(spec string, cmd func()) (interface{}, error) {
	return s.Cron.AddFunc(spec, cmd)
}

func (s *CronScheduler) Start() {
	s.Cron.Start()
}

var _ IScheduler = (*CronScheduler)(nil) // Ensure CronScheduler implements Scheduler
