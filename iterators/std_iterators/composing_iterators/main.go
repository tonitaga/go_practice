package main

import (
	"fmt"
	"iter"
)

type IntegerSequence = iter.Seq[int]

func Integers(size int) IntegerSequence {
	return func(yield func(int) bool) {
		for value := range size {
			if !yield(value) {
				break
			}
		}
	}
}

func Even(sequence IntegerSequence) IntegerSequence {
	return func(yield func(int) bool) {
		for value := range sequence {
			if value%2 != 0 {
				continue
			}

			if !yield(value) {
				break
			}
		}
	}
}

func main() {
	for value := range Even(Integers(100)) {
		fmt.Println(value)
	}
}
