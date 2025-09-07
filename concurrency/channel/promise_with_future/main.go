package main

import (
	"fmt"
	"time"
)

type Future[T any] struct {
	resultChan <-chan T
}

func NewFuture[T any](resultChan <-chan T) Future[T] {
	return Future[T]{
		resultChan: resultChan,
	}
}

func (f *Future[T]) Get() T {
	return <-f.resultChan
}

type Promise[T any] struct {
	futureChan chan T
}

func (p *Promise[T]) Set(value T) {
	p.futureChan <- value
	close(p.futureChan)
}

func NewPromise[T any]() Promise[T] {
	return Promise[T]{
		futureChan: make(chan T),
	}
}

func (p *Promise[T]) GetFuture() Future[T] {
	return NewFuture(p.futureChan)
}

func main() {
	promise := NewPromise[string]()
	future := promise.GetFuture()

	go func() {
		time.Sleep(time.Second)
		promise.Set("Hello, world")
	}()

	fmt.Println(future.Get())
}
