package main

import "fmt"

func NewIterator(size int) func(func(int) bool) {
	return func(yield func(int) bool) {
		for value := range size {
			if !yield(value) {
				return
			}
		}
	}
}

func main() {
	for value := range NewIterator(10) {
		fmt.Println(value)
	}
}
