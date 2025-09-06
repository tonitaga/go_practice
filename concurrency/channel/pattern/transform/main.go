package main

import (
	"fmt"
	"sync"
)

/*
Transform не устоявшееся название. Трансформирует каждое значение в readonly канале
*/

type TransformOperation[T any] = func(T) T

func Transform[T any](inputChan <- chan T, unaryOp TransformOperation[T]) <- chan T {
	outputChan := make(chan T)

	go func ()  {
		defer close(outputChan)

		for value := range inputChan {
			outputChan <- unaryOp(value)
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

	squareOperation := func (value int) int {
		return value * value;
	}

	transformedChan := Transform(mainChan, squareOperation)

	wg := sync.WaitGroup{}
	wg.Go(func ()  {
		for value := range transformedChan {
			fmt.Println(value)
		}
	})

	wg.Wait()
}