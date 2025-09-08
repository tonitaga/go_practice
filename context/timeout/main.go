package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	timeoutContext, cancelContext := context.WithTimeout(context.Background(), 1*time.Second)

	defer cancelContext()

	done := make(chan struct{})

	go func() {
		defer close(done)
		time.Sleep(time.Second * 2)
	}()

	select {
	case <-done:
		fmt.Println("Task done")
	case <-timeoutContext.Done():
		fmt.Println("Timeout")
	}
}
