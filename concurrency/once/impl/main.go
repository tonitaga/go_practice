package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type CallOnce struct {
	done  atomic.Bool
	mutex sync.Mutex
}

func (c *CallOnce) Once(action func()) {
	if c.done.Load() {
		return
	}

	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.done.Load() {
		return
	}

	defer c.done.Store(true)
	action()
}

func main() {
	callOnce := CallOnce{}

	wg := sync.WaitGroup{}
	for i := 0; i != 5; i++ {
		wg.Go(func() {
			callOnce.Once(func() {
				time.Sleep(time.Second)
				fmt.Printf("[Go #%d] Called once\n", i)
			})

			fmt.Printf("[Go #%d] End\n", i)
		})
	}

	wg.Wait()
}
