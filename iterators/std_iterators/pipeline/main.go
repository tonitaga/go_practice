package main

import (
	"fmt"
	"iter"
)

func MakeIntegerSequence(size int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for value := range size {
			if !yield(value) {
				break
			}
		}
	}
}

func Multiply(sequence iter.Seq[int], number int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for value := range sequence {
			if !yield(value * number) {
				break
			}
		}
	}
}

func Filter(sequence iter.Seq[int], filter func(int) bool) iter.Seq[int] {
	return func(yield func(int) bool) {
		for value := range sequence {
			if !filter(value) {
				continue
			}

			if !yield(value) {
				break
			}
		}
	}
}

func main() {
	sequence := MakeIntegerSequence(100)
	sequence = Multiply(sequence, 3)
	sequence = Filter(sequence, func(i int) bool { return i%2 == 1 })

	for value := range sequence {
		fmt.Println(value)
	}
}
