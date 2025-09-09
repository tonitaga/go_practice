package main

import (
	"fmt"
	"time"
)

func main() {
	scheduler := New()
	scheduler.AddPeriodic(func() bool {
		fmt.Println("Println")
		return false
	}, time.Second)

	go func() {
		time.Sleep(time.Second * 5)
		scheduler.Stop()
	}()

	scheduler.Run()
}
