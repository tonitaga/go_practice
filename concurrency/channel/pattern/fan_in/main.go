package main

import (
	"fmt"
	"sync"
)

/*
Fan-in — паттерн многопоточности в Go (Golang), который объединяет несколько входных каналов или потоков данных в один выходной канал.
Это позволяет консолидировать результаты из различных источников.
*/

func MergeChannels[T any](channels ...<-chan T) <-chan T {
	wg := sync.WaitGroup{}

	outChan := make(chan T)
	for _, channel := range channels {
		wg.Go(func ()  {
			for item := range channel {
				outChan <- item
			}
		})
	}

	go func ()  {
		wg.Wait()
		close(outChan)
	}()

	return outChan
}

func main() {
	channel1 := make(chan int)
	channel2 := make(chan int)
	channel3 := make(chan int)

	wg := sync.WaitGroup{};
	wg.Go(func ()  {
		defer func ()  {
			close(channel1)
			close(channel2)
			close(channel3)
		}()

		count := 16
		for i := 0; i <= count; i++ {
			channel1 <- count * 1 + i
			channel2 <- count * 2 + i
			channel3 <- count * 3 + i
		}
	})

	mergedChannels := MergeChannels(channel1, channel2, channel3)
	for value := range mergedChannels {
		fmt.Println(value)
	}

	wg.Wait()
}