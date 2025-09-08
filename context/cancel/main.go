package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	cancelContext, cancelFunc := context.WithCancel(context.Background())

	go func() {
		defer cancelFunc()
		time.Sleep(time.Second * 2)
	}()

	<-cancelContext.Done()
	fmt.Println("Cancel done")
}
