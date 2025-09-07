package main

import (
	"fmt"
)

type FutureStatus int

type Future[T any] struct {
	resultChan chan T
}

func NewFuture[T any](asyncOperation func() T) *Future[T] {
	future := &Future[T]{
		resultChan: make(chan T),
	}

	go func() {
		future.resultChan <- asyncOperation()
		close(future.resultChan)
	}()

	return future
}

func (f *Future[T]) Get() T {
	return <-f.resultChan
}

func main() {
	asyncOperation := func() string {
		return "Hello, world"
	}

	future := NewFuture(asyncOperation)
	fmt.Println(future.Get())
}
