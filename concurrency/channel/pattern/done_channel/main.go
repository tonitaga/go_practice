package main

import (
	"fmt"
	"time"
)

func Process(stop <- chan bool) <- chan bool {
	done := make(chan bool)

	go func() {
		defer close(done)

		ticker := time.NewTicker(1000 * time.Millisecond)

		for {
			select {
			case <- stop:
				fmt.Println("--- Stopping ---")
				return
			case <- ticker.C:
				fmt.Println("--- Processing ---")
			}
		}
	}()

	return done
}

func main() {
	stop := make(chan bool)
	done := Process(stop)

	time.Sleep(time.Second * 4)

	close(stop)
	<-done

	fmt.Println("--- Stopped ---")
}