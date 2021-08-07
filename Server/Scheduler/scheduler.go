package Scheduler

import (
	"context"
	"sync"
	"time"
)

type Scheduler struct {
	wg            *sync.WaitGroup
	cancellations []context.CancelFunc
}

func NewScheduler() *Scheduler {
	return &Scheduler{
		wg:            new(sync.WaitGroup),
		cancellations: make([]context.CancelFunc, 0),
	}
}

type Job func(context.Context)

func (scheduler *Scheduler) runJob(ctx context.Context, job Job, interval time.Duration) {
	ticker := time.NewTicker(interval)
	
	for {
		select {
		case <-ticker.C:
			job(ctx)
		case <-ctx.Done():
			scheduler.wg.Done()
			return
		}
	}
}

func (scheduler *Scheduler) Add(ctx context.Context, job Job, interval time.Duration) {
	ctx, cancel := context.WithCancel(ctx)
	scheduler.cancellations = append(scheduler.cancellations, cancel)

	scheduler.wg.Add(1)
	go scheduler.runJob(ctx, job, interval)
}
