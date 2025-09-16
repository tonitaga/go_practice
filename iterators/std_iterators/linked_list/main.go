package main

import (
	"fmt"
	"iter"
)

type LinkedNode[T any] struct {
	value T
	next  *LinkedNode[T]
}

type LinkedList[T any] struct {
	head *LinkedNode[T]
}

func (l *LinkedList[T]) Push(value T) {
	l.head = &LinkedNode[T]{value: value, next: l.head}
}

func (l *LinkedList[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		for node := l.head; node != nil; node = node.next {
			if !yield(node.value) {
				break
			}
		}
	}
}

func main() {
	list := &LinkedList[string]{}

	for i := range 10 {
		list.Push(fmt.Sprintf("Data #%d", i))
	}

	for value := range list.All() {
		fmt.Println(value)
	}
}
