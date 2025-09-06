package main

import (
	"fmt"
	"time"
)

type Semaphore struct {
	value chan int
}

func NewSemaphore(initialValue int) *Semaphore {
	return &Semaphore{
		value: make(chan int, initialValue),
	}
}

func (s *Semaphore) Acquire() {
	s.value <- 0
}

func (s *Semaphore) Release() {
	<-s.value
}

func (s *Semaphore) Go(routine func()) {
	s.Acquire()
	go func() {
		defer s.Release()
		routine()
	}()
}

func main() {
	semaphore := NewSemaphore(2)

	semaphore.Go(func() {
		fmt.Println("1")
		time.Sleep(time.Second * 5)
	})

	semaphore.Go(func() {
		fmt.Println("2")
		time.Sleep(time.Second * 5)
	})

	// Must wait release
	semaphore.Go(func() {
		fmt.Println("3")
		time.Sleep(time.Second * 5)
	})

	time.Sleep(time.Second)
}
