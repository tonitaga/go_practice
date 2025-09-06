package main

import (
	"fmt"
	"sync"
)

/*
Tee-паттерн в Go (golang) — это паттерн конкурентного программирования,
который принимает один входной канал и дублирует его данные в несколько выходных каналов.
*/

func TeeChannels[T any](inputChan <-chan T, n int) []<-chan T{
	outputChans := make([]chan T, 0, n)
	for i := 0; i < n; i++ {
		outputChans = append(outputChans, make(chan T))
	}

	go func() {
		defer func ()  {
			for _, channel := range outputChans {
				close(channel)
			}
		}()

		for value := range inputChan {
			for _, channel := range outputChans {
				channel <- value
			}
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
		
		for i := 0; i != 4; i++ {
			mainChan <- i;
		}
	}()

	channelsCount := 3
	channels := TeeChannels(mainChan, channelsCount);

	wg := sync.WaitGroup{}

	for i := range channelsCount {
		wg.Go(func ()  {
			for value := range channels[i] {
				fmt.Printf("chan[%d] %d\n", i, value)
			}
		})
	}

	wg.Wait()
}