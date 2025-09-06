package main

import (
	"fmt"
	"time"
)

type UntilDoneOrStopWorker[T any] struct {
	C         chan T
	stopToken chan bool
}

func NewWorker[T any](dataChannel <-chan T) *UntilDoneOrStopWorker[T] {
	worker := &UntilDoneOrStopWorker[T]{}

	worker.C = make(chan T)
	worker.stopToken = make(chan bool)

	go func() {
		defer close(worker.C)

		for {
			/*
				Дополнительная приоритезация для того,
				чтобы избежать случайного выбора ветки в момент одновременного получения
				сигнала на стоп и чтения значения
			*/
			select {
			case <-worker.stopToken:
				return
			default:
			}

			select {
			case value, opened := <-dataChannel:
				if !opened {
					return
				}

				worker.C <- value
			case <-worker.stopToken:
				return
			}
		}
	}()

	return worker
}

func (w *UntilDoneOrStopWorker[T]) Shutdown() {
	close(w.stopToken)
}

func main() {
	dataChannel := make(chan string)

	worker := NewWorker(dataChannel)

	go func() {
		for i := 0; i < 500; i++ {
			dataChannel <- fmt.Sprintf("Data #%d", i)
			time.Sleep(time.Millisecond * 100)
		}

		close(dataChannel)
	}()

	go func() {
		time.Sleep(time.Second)
		worker.Shutdown()
	}()

	for value := range worker.C {
		fmt.Println(value)
	}
}
