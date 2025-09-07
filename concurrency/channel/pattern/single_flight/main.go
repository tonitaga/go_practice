package main

import (
	"fmt"
	"sync"
	"time"
)

/*
Основная цель Single Flight — гарантировать, что в любой момент времени выполняется только
один вызов дорогостоящей или дублирующей операции.
Когда несколько горутин запрашивают один и тот же ресурс, Single Flight гарантирует,
что функция выполняется только один раз, а результат делится между всеми запрашивающими.
*/

type call[T any] struct {
	err      error
	value    T
	doneChan chan struct{}
}

type SingleFlight[T comparable, Y any] struct {
	mutex sync.Mutex
	calls map[T]*call[Y]
}

func NewSingleFlight[T comparable, Y any]() *SingleFlight[T, Y] {
	return &SingleFlight[T, Y]{
		mutex: sync.Mutex{},
		calls: make(map[T]*call[Y]),
	}
}

func (singleFlight *SingleFlight[T, Y]) Do(key T, action func() (Y, error)) (Y, error) {
	singleFlight.mutex.Lock()
	if call, found := singleFlight.calls[key]; found {
		singleFlight.mutex.Unlock()
		return singleFlight.wait(call)
	}

	call := &call[Y]{
		doneChan: make(chan struct{}),
	}

	singleFlight.calls[key] = call
	singleFlight.mutex.Unlock()

	go func() {
		defer func() {
			singleFlight.mutex.Lock()

			close(call.doneChan)
			delete(singleFlight.calls, key)

			singleFlight.mutex.Unlock()
		}()

		call.value, call.err = action()
	}()

	return singleFlight.wait(call)
}

func (singleFlight *SingleFlight[T, Y]) wait(call *call[Y]) (Y, error) {
	<-call.doneChan
	return call.value, call.err
}

type User struct {
	Username string
	Email    string
	Age      uint
}

func main() {
	singleFlight := NewSingleFlight[string, User]()

	fmt.Println("=== started ===")

	wg := sync.WaitGroup{}
	for i := 0; i != 5; i++ {
		wg.Go(func() {
			user, err := singleFlight.Do("hello_world", func() (User, error) {
				time.Sleep(time.Second * 2) // Hard work for getting data
				return User{
					Username: "tonitaga",
					Email:    "gubaydullin.nurislam@gmail.com",
					Age:      22,
				}, nil
			})

			if err != nil {
				fmt.Println("[", i, "]", err)
				return
			}

			fmt.Println("[", i, "]", user)
		})
	}

	wg.Wait()
}
