package main

import (
	"fmt"
	"sync"
)

type FilterOperation[T any] = func(T) bool

func FilterChannel[T any](inputChan <- chan T, filterOp FilterOperation[T]) <- chan T {
	outputChan := make(chan T)

	go func ()  {
		defer close(outputChan)

		for value := range inputChan {
			if (filterOp(value)) {
				outputChan <- value
			}
		}
	}()

	return outputChan
} 

func main() {
	mainChan := make(chan int)

	go func ()  {
		defer close(mainChan)

		for i := 0; i < 8; i++ {
			mainChan <- i
		}
	}()

	isOdd := func (value int) bool {
		return value % 2 == 0;
	}

	filteredChan := FilterChannel(mainChan, isOdd)

	wg := sync.WaitGroup{}
	wg.Go(func ()  {
		for value := range filteredChan {
			fmt.Println(value)
		}
	})

	wg.Wait()
}