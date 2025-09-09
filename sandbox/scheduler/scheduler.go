package main

import (
	"math"
	"sync"
	"time"
)

type SchedulerTaskFunc = func() bool

type SchedulerTask struct {
	Task     SchedulerTaskFunc
	Period   time.Duration
	NextCall time.Time
}

type Scheduler struct {
	tasks         map[int]SchedulerTask
	mutex         *sync.Mutex
	closeChan     chan struct{}
	closeDoneChan chan struct{}
	taskAdd       chan struct{}
	lastID        int
}

func New() *Scheduler {
	m := &sync.Mutex{}

	return &Scheduler{
		tasks:         make(map[int]SchedulerTask),
		mutex:         m,
		closeChan:     make(chan struct{}),
		closeDoneChan: make(chan struct{}),
		taskAdd:       make(chan struct{}),
		lastID:        0,
	}
}

func (s *Scheduler) AddOneShot(task SchedulerTaskFunc) {
	s.AddPeriodic(task, 0)
}

func (s *Scheduler) AddPeriodic(task SchedulerTaskFunc, period time.Duration) {
	{
		defer s.mutex.Unlock()

		s.mutex.Lock()
		s.tasks[s.lastID] = SchedulerTask{
			Task:     task,
			Period:   period,
			NextCall: time.Now().Add(period),
		}

		s.lastID += 1
	}

	s.taskAdd <- struct{}{}
}

func (s *Scheduler) Stop() {
	close(s.closeChan)
	close(s.taskAdd)
	<-s.closeDoneChan
}

func (s *Scheduler) Run() {
	defer close(s.closeDoneChan)

	for {
		s.mutex.Lock()
		if len(s.tasks) == 0 {
			<-s.taskAdd
		}

		select {
		case <-s.closeChan:
			return
		default:
		}

		now := time.Now()

		minCallTime := time.Duration(math.MaxInt64)
		minCallTimeTaskID := -1

		var minCallTimeTask SchedulerTask

		for id, task := range s.tasks {
			callTime := task.NextCall.Sub(now)
			if callTime.Nanoseconds() < minCallTime.Nanoseconds() {
				// Возможно просрочили задачу, поэтому чтобы не было отрицательного
				minCallTime = max(callTime, time.Duration(0))
				minCallTimeTaskID = id
			}
		}

		if minCallTimeTaskID == -1 {
			panic("Launching task should be available there")
		}

		minCallTimeTask = s.tasks[minCallTimeTaskID]

		s.mutex.Unlock()

		ticker := time.NewTicker(minCallTime)

		select {
		case <-s.closeChan:
			return
		case <-ticker.C:
			needRemove := minCallTimeTask.Task()
			if needRemove || minCallTimeTask.Period == 0 {
				defer s.mutex.Unlock()

				s.mutex.Lock()
				delete(s.tasks, minCallTimeTaskID)
			}
			continue
		case <-s.taskAdd:
			continue
		}
	}
}
