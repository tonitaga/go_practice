package main

import (
	"fmt"
	"sync"
	"time"
)

type BigData string

func FunOutChannel[T any](inputChan <-chan T, count int) []<-chan T {
	outputChannels := make([]chan T, 0, count)
	for i := 0; i != count; i++ {
		outputChannels = append(outputChannels, make(chan T))
	}

	go func() {
		defer func() {
			for _, channel := range outputChannels {
				close(channel)
			}
		}()

		current := 0
		for value := range inputChan {
			outputChannels[current] <- value
			current = (current + 1) % count
		}
	}()

	readonlyChannels := make([]<-chan T, 0, count)
	for i := 0; i != count; i++ {
		readonlyChannels = append(readonlyChannels, outputChannels[i])
	}

	return readonlyChannels
}

const WorkersCount int = 2

func main() {
	mainChan := make(chan BigData)

	begin := time.Now()

	outputChannels := FunOutChannel(mainChan, WorkersCount)

	wg := sync.WaitGroup{}
	for i := 0; i != WorkersCount; i++ {
		dataChan := outputChannels[i]
		wg.Go(func() {
			resultValue := BigData("")
			for value := range dataChan {
				resultValue = resultValue + value
			}

			// fmt.Printf("Go #%d handles data '%s'\n", i, resultValue)
		})
	}

	wg.Go(func() {
		defer close(mainChan)

		for i := 0; i != WorkersCount*1024; i++ {
			data := fmt.Sprintf(
				`{
	"message_id": %d,
	"name": "tonitaga_%d",
}`, i, i)

			mainChan <- BigData(data)
		}
	})

	wg.Wait()

	fmt.Println("Handling data time (us):", float64(time.Since(begin).Microseconds())/1000.0)
}
