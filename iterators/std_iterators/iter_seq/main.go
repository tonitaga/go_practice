package main

import (
	"fmt"
	"iter"
)

func MakeIterOne(size int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for value := range size {
			if !yield(value) {
				return
			}
		}
	}
}

func MakeIterTwo(array []string) iter.Seq2[int, string] {
	return func(yield func(int, string) bool) {
		for i, value := range array {
			if !yield(i, value) {
				return
			}
		}
	}
}

func main() {
	for value := range MakeIterOne(10) {
		fmt.Println(value)
	}

	for i, value := range MakeIterTwo([]string{"H", "E", "L", "L", "O"}) {
		fmt.Println(i, value)
	}
}
