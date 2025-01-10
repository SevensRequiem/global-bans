package schedule

import (
	"fmt"
	"globalbans/backend/logs"
	"time"
)

type Task struct {
	Action   func()
	Duration time.Duration
}

type Scheduler struct {
	StartTime          time.Time
	LastUpdate         time.Time
	LastUpdateDuration time.Duration
	Tasks              []Task
}

func NewScheduler() *Scheduler {
	return &Scheduler{
		StartTime: time.Now(),
		Tasks:     []Task{},
	}
}

func (s *Scheduler) ScheduleTask(task Task) {
	s.Tasks = append(s.Tasks, task)
	logs.LogInfo(fmt.Sprintf("Task scheduled for every %v", task.Duration), 0, "scheduler/scheduler.go")
}

func (s *Scheduler) Run() {
	s.StartTime = time.Now()
	logs.LogInfo("Scheduler started", 0, "scheduler/scheduler.go")

	for _, task := range s.Tasks {
		go func(t Task) {
			ticker := time.NewTicker(t.Duration)
			defer ticker.Stop()

			for range ticker.C {
				t.Action()
				s.LastUpdate = time.Now()
				s.LastUpdateDuration = s.LastUpdate.Sub(s.StartTime)
				logs.LogInfo(fmt.Sprintf("Task executed at %v", s.LastUpdate), 0, "scheduler/scheduler.go")
			}
		}(task)
	}
}
