package main

import (
	"fmt"
	"time"
)

type Worker struct {
	stopToken   chan bool
	doneToken   chan bool
	userRoutine func()
	period      time.Duration
	isStopped   bool
}

func NewWorker(period time.Duration, routine func()) *Worker {
	return &Worker{
		stopToken:   nil,
		doneToken:   nil,
		userRoutine: routine,
		period:      period,
		isStopped:   true,
	}
}

func (w *Worker) Launch() {
	if !w.isStopped {
		panic("worker already working")
	}

	w.isStopped = false

	w.stopToken = make(chan bool)
	w.doneToken = make(chan bool)

	go func() {
		defer close(w.doneToken)

		ticker := time.NewTicker(w.period)

		for {
			/*
				Дополнительная приоритезация для того,
				чтобы избежать случайного выбора ветки в момент одновременного получения
				сигнала на стоп и истечению тикера
			*/
			select {
			case <-w.stopToken:
				return
			default:
			}

			select {
			case <-w.stopToken:
				return
			case <-ticker.C:
				w.userRoutine()
			}
		}

	}()
}

func (w *Worker) Shutdown() {
	if w.isStopped {
		panic("worker already stopped")
	}

	w.isStopped = true

	close(w.stopToken)
	<-w.doneToken
}

func main() {
	worker := NewWorker(time.Millisecond*500, func() {
		fmt.Println("=== Processing ===")
	})

	worker.Launch()

	time.Sleep(time.Second * 3)

	worker.Shutdown()

}
