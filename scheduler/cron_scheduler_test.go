package scheduler

import (
	"testing"
	"time"
)

func TestAddFuncAndStart(t *testing.T) {
	scheduler := NewCronScheduler()

	executed := false
	scheduler.AddFunc("@every 1s", func() {
		executed = true
	})

	scheduler.Start()
	time.Sleep(2 * time.Second)

	if !executed {
		t.Errorf("Expected the scheduled function to be executed")
	}
}
