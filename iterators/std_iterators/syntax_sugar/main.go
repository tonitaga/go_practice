package main

import (
	"fmt"
	"iter"
)

func FilterRange(sequence []int, filter func(int) bool) iter.Seq[int] {
	return func(yield func(int) bool) {
		for _, value := range sequence {
			if filter(value) {
				if !yield(value) {
					return
				}
			}
		}
	}
}

func main() {
	iterator := FilterRange([]int{1, 2, 3, 4}, func(v int) bool { return v%2 == 0 })

	for v := range iterator {
		fmt.Println(v)
		if v == 5 {
			break
		}
	}

	iterator(func(v int) bool {
		fmt.Println(v)
		if v == 5 {
			return false // like break
		}

		return true // like continue or end of iteration
	})
}
