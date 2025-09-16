package main

import "fmt"

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

func (l *LinkedList[T]) GetIterator() *LinkedListIterator[T] {
	return &LinkedListIterator[T]{
		current: l.head,
	}
}

type LinkedListIterator[T any] struct {
	current *LinkedNode[T]
}

func (it *LinkedListIterator[T]) HasNext() bool {
	return it.current != nil
}

func (it *LinkedListIterator[T]) GetNext() *LinkedNode[T] {
	node := it.current
	it.current = it.current.next
	return node
}

func main() {
	list := &LinkedList[int]{}

	for i := range 10 {
		list.Push(i)
	}

	iterator := list.GetIterator()
	for iterator.HasNext() {
		node := iterator.GetNext()
		fmt.Println(*node)
	}
}
