package main

import (
	"fmt"
	"time"
)

type promiseResult[T any] struct {
	value T
	err   error
}

type Promise[T any] struct {
	resultChannel chan promiseResult[T]
}

func NewPromise[T any](asyncRoutine func() (T, error)) Promise[T] {
	promise := Promise[T]{
		resultChannel: make(chan promiseResult[T]),
	}

	go func() {
		defer close(promise.resultChannel)

		value, err := asyncRoutine()

		promise.resultChannel <- promiseResult[T]{value: value, err: err}
	}()

	return promise
}

func (p *Promise[T]) Then(successCallback func(result T), errorCallback func(err error)) {
	go func() {
		result := <-p.resultChannel

		if result.err != nil {
			errorCallback(result.err)
			return
		}

		successCallback(result.value)
	}()
}

func asyncSuccessOperation() (string, error) {
	return "ok", nil
}

func asyncFailedOperation() (string, error) {
	return "", fmt.Errorf("some error")
}

func main() {
	promise1 := NewPromise(asyncFailedOperation)
	promise1.Then(
		func(value string) {
			fmt.Println(value)
		},
		func(err error) {
			fmt.Printf("Got error: %s\n", err)
		},
	)

	promise2 := NewPromise(asyncSuccessOperation)
	promise2.Then(
		func(value string) {
			fmt.Println(value)
		},
		func(err error) {
			fmt.Printf("Got error: %s\n", err)
		},
	)

	time.Sleep(time.Second)
}
