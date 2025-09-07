package main

import "fmt"

func Generate(left, right int) <-chan int {
	resultChan := make(chan int)

	go func() {
		defer close(resultChan)

		for i := left; i < right; i++ {
			resultChan <- i
		}
	}()

	return resultChan
}

func main() {
	for value := range Generate(1, 8) {
		fmt.Println(value)
	}
}
