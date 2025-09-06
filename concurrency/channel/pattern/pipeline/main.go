package main

import "fmt"

func square(value int) int {
	return value * value
}

func GenerateChannel[T any](values ...T) <- chan T {
	outputChan := make(chan T)

	go func ()  {
		defer close(outputChan)

		for _, value := range values {
			outputChan <- value
		}
	}()

	return outputChan
}

func ProcessChannel[T any](inputChan <- chan T, action func(T) T) <- chan T {
	outputChan := make(chan T)
	
	go func() {
		defer close(outputChan)

		for value := range inputChan {
			outputChan <- action(value)
		}
	}()

	return outputChan
}

func main() {
	values := []int {0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	for value := range ProcessChannel(GenerateChannel(values...), square) {
		fmt.Println(value)
	}
}