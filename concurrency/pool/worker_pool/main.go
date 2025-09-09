package main

import (
	"log"
	"sync"
)

type WorkerPool struct {
	tasksChan chan func()

	closeChan     chan struct{}
	closeDoneChan chan struct{}
	wg            sync.WaitGroup
}

func (w *WorkerPool) launchWorkers(workersCount int) {
	for i := 0; i != workersCount; i++ {
		w.wg.Go(func() {
			w.workerRoutine(i)
		})
	}
}

func (w *WorkerPool) workerRoutine(workerIndex int) {
	defer log.Printf("[worker #%d] stopped\n", workerIndex)

	for {
		select {
		case <-w.closeChan:
			return
		case task, ok := <-w.tasksChan:
			if !ok {
				log.Printf("[worker #%d] exited in <-w.tasksChan\n", workerIndex)
				return
			}

			task()
			log.Printf("[worker #%d] made task\n", workerIndex)
		}
	}
}

func (w *WorkerPool) Do(task func()) {
	select {
	case <-w.closeChan:
		log.Println("Worker pool is closed, cannot accept new tasks")
		return
	default:
		w.tasksChan <- task
	}
}

func (w *WorkerPool) GracefulShutdown() {
	close(w.tasksChan)
	w.wg.Wait()
	close(w.closeDoneChan)
	close(w.closeChan)
}

func (w *WorkerPool) ForceShutdown() {
	close(w.closeChan)
	w.wg.Wait()
	close(w.tasksChan)
	close(w.closeDoneChan)
}

func NewWorkerPool(workersCount int, tasksCapacity int) *WorkerPool {
	pool := &WorkerPool{
		tasksChan:     make(chan func(), tasksCapacity),
		closeChan:     make(chan struct{}),
		closeDoneChan: make(chan struct{}),
	}

	pool.launchWorkers(workersCount)
	return pool
}

func main() {
	workerPool := NewWorkerPool(4, 5000)

	for i := 0; i < 2000; i++ {
		workerPool.Do(func() {
			log.Println("Value:", i)
		})
	}

	workerPool.GracefulShutdown()
}
