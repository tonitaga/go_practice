package main

import (
	"fmt"
	"iter"
	"slices"
)

func Backward[T any](data []T) iter.Seq[T] {
	return func(yield func(T) bool) {
		for i := len(data) - 1; i >= 0; i-- {
			if !yield(data[i]) {
				break
			}
		}
	}
}

func main() {
	data := []int{2, 4, 5, 64, 8}

	for value := range data {
		fmt.Println(value)
	}

	fmt.Println("===")

	for value := range Backward(data) {
		fmt.Println(value)
	}

	fmt.Println("===")

	for _, value := range slices.Backward(data) {
		fmt.Println(value)
	}
}
