package main

import (
	"fmt"
	"sync"
)

/*
Fan-out — паттерн конкурентного программирования в Go (Golang),
который распределяет задачи между несколькими горутинами. Это позволяет:
*/

func SplitChannels[T any](inputChan <-chan T, n int) []<-chan T {
	outputChans := make([]chan T, /*size=*/0, /*capacity*/n)
	for i := 0; i < n; i++ {
		outputChans = append(outputChans, make(chan T))
	}

	go func() {
		defer func ()  {
			for _, channel := range outputChans {
				close(channel)
			}
		}()

		index := 0
		for value := range inputChan {
			outputChans[index] <- value
			index = (index + 1) % n // Round Robin (круговой перебор, круговая ротация)
		}
	}()

	resultChans := make([]<-chan T, n)
	for i, channel := range outputChans {
		resultChans[i] = channel
	}

	return resultChans
}

func main() {
	mainChan := make(chan int)

	go func ()  {
		defer close(mainChan)

		for i := 0; i < 32; i++ {
			mainChan <- i
		}
	}()

	channels := SplitChannels(mainChan, 2)

	wg := sync.WaitGroup{}
	
	wg.Go(func ()  {
		for value := range channels[0] {
			fmt.Printf("chan[0]: %d\n", value)
		}
	})

	wg.Go(func ()  {
		for value := range channels[1] {
			fmt.Printf("chan[1]: %d\n", value)
		}
	})

	wg.Wait()
}
