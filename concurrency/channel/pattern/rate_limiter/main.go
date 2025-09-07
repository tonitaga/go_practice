package main

import (
	"fmt"
	"sync"
	"time"
)

type RateLimiter struct {
	bufferedChan  chan struct{}
	closeDone     chan struct{}
	closeDoneChan chan struct{}
}

func NewRateLimiter(limit int64, period time.Duration) RateLimiter {
	rateLimiter := RateLimiter{
		bufferedChan:  make(chan struct{}, 1),
		closeDone:     make(chan struct{}),
		closeDoneChan: make(chan struct{}),
	}

	go func() {
		interval := period.Nanoseconds() / limit
		rateLimiter.runPeriodicTask(time.Duration(interval))
	}()

	return rateLimiter
}

func (rateLimiter *RateLimiter) runPeriodicTask(interval time.Duration) {
	ticker := time.NewTicker(interval)

	defer func() {
		ticker.Stop()
		close(rateLimiter.closeDoneChan)
	}()

	for {
		select {
		case <-rateLimiter.closeDone:
			return
		default:
		}

		select {
		case <-rateLimiter.closeDone:
			return
		case <-ticker.C:
			select {
			case <-rateLimiter.bufferedChan:
			default:
			}
		}
	}
}

func (rateLimiter *RateLimiter) AllowCall() bool {
	select {
	case rateLimiter.bufferedChan <- struct{}{}:
		return true
	default:
		return false
	}
}

func (rateLimiter *RateLimiter) Stop() {
	close(rateLimiter.closeDone)
	<-rateLimiter.closeDoneChan
}

func main() {
	fmt.Println("=== started ===")

	rateLimiter := NewRateLimiter(10, time.Second)

	action := func(actionIndex int) {
		fmt.Println("[", time.Now().Format(time.StampMilli), "]", "[", actionIndex, "] Done")
	}

	wg := sync.WaitGroup{}
	wg.Go(func() {
		for i := 0; i < 100; i++ {
			time.Sleep(time.Millisecond * 10)
			if rateLimiter.AllowCall() {
				action(i)
			}
		}
	})

	wg.Wait()
	rateLimiter.Stop()
}
