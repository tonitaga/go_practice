package main

import (
	"fmt"
	"iter"
)

type Iterator[T any] struct {
	iterator iter.Seq[T]
}

func (it *Iterator[T]) Collect() []T {
	data := make([]T, 0)
	for value := range it.iterator {
		data = append(data, value)
	}

	return data
}

func For[T any](data []T) *Iterator[T] {
	return &Iterator[T]{
		iterator: func(yield func(T) bool) {
			for _, value := range data {
				if !yield(value) {
					break
				}
			}
		},
	}
}

func (it *Iterator[T]) Map(action func(T) T) *Iterator[T] {
	iteratorCopy := it.iterator
	it.iterator = func(yield func(T) bool) {
		for value := range iteratorCopy {
			value = action(value)
			if !yield(value) {
				return
			}
		}
	}

	return it
}

func (it *Iterator[T]) Filter(filter func(T) bool) *Iterator[T] {
	iteratorCopy := it.iterator
	it.iterator = func(yield func(T) bool) {
		for value := range iteratorCopy {
			if !filter(value) {
				continue
			}

			if !yield(value) {
				return
			}
		}
	}

	return it
}

func (it *Iterator[T]) Reverse() *Iterator[T] {
	data := it.Collect()

	currentIndex := len(data) - 1
	for value := range it.iterator {
		data[currentIndex] = value
		currentIndex--
	}

	return For(data)
}

func (it *Iterator[T]) Each(action func(T)) *Iterator[T] {
	for value := range it.iterator {
		action(value)
	}

	return it
}

func Square(value int) int {
	return value * value
}

func IsEven(value int) bool {
	return value%2 == 0
}

func Print(value int) {
	fmt.Printf("value: %v\n", value)
}

func main() {
	For([]int{1, 2, 3, 4, 5, 6, 7, 8}).
		Map(Square).
		Each(Print).
		Filter(IsEven).
		Each(Print).
		Reverse().
		Each(Print)
}
