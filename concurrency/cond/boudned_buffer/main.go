package main

import (
	"fmt"
	"sync"
	"time"
)

type BlockingQueue[T any] struct {
	buffer   []T
	mutex    *sync.Mutex
	notEmpty *sync.Cond
}

func NewBlockingQueue[T any]() BlockingQueue[T] {
	queue := BlockingQueue[T]{
		buffer: make([]T, 0),
		mutex:  &sync.Mutex{},
	}

	queue.notEmpty = sync.NewCond(queue.mutex)
	return queue
}

func (this *BlockingQueue[T]) Push(value T) {
	defer this.mutex.Unlock()

	this.mutex.Lock()
	this.buffer = append(this.buffer, value)
	this.notEmpty.Signal()
}

func (this *BlockingQueue[T]) Front() T {
	defer this.mutex.Unlock()

	this.mutex.Lock()
	if len(this.buffer) == 0 {
		this.notEmpty.Wait()
	}

	// Deadlock

	value := this.buffer[0]

	this.buffer = this.buffer[1:]
	return value
}

func main() {
	buffer := NewBlockingQueue[string]()

	wg := sync.WaitGroup{}

	doneChan := make(chan bool)

	wg.Go(func() {
		defer close(doneChan)
		for i := 0; i != 20; i++ {
			buffer.Push(fmt.Sprintf("[Data] #%d", i))
			time.Sleep(time.Millisecond * 200)
		}
	})

	wg.Go(func() {
		for {
			select {
			case <-doneChan:
				return
			default:
				fmt.Println("[Go #1] Got:", buffer.Front())
			}
		}
	})

	wg.Wait()
}
