package main

import (
	"fmt"
	"iter"
)

type TreeNode[T any] struct {
	value T
	left  *TreeNode[T]
	right *TreeNode[T]
}

func (t *TreeNode[T]) PreOrder() iter.Seq[T] {
	return func(yield func(T) bool) {
		t.preOrderHelper(yield)
	}
}

func (t *TreeNode[T]) InOrder() iter.Seq[T] {
	return func(yield func(T) bool) {
		t.inOrderHelper(yield)
	}
}

func (t *TreeNode[T]) preOrderHelper(yield func(T) bool) bool {
	if t == nil {
		return true
	}

	return yield(t.value) && t.left.preOrderHelper(yield) && t.right.preOrderHelper(yield)
}

func (t *TreeNode[T]) inOrderHelper(yield func(T) bool) bool {
	if t == nil {
		return true
	}

	return t.left.inOrderHelper(yield) && yield(t.value) && t.right.inOrderHelper(yield)
}

func main() {
	/*
					5
			3				7
		2		4		6		8
	*/

	tree := &TreeNode[int]{
		value: 5,
		left: &TreeNode[int]{
			value: 3,
			left: &TreeNode[int]{
				value: 2,
			},
			right: &TreeNode[int]{
				value: 4,
			},
		},
		right: &TreeNode[int]{
			value: 7,
			left: &TreeNode[int]{
				value: 6,
			},
			right: &TreeNode[int]{
				value: 8,
			},
		},
	}

	fmt.Println("PreOrder:")
	for value := range tree.PreOrder() {
		fmt.Println(value)
	}

	fmt.Println("InOrder:")
	for value := range tree.InOrder() {
		fmt.Println(value)
	}
}
