package main

import (
	"fmt"
	"iter"
	"slices"
)

func main() {
	values := slices.Values([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})

	next, stop := iter.Pull(values)
	defer stop()

	for {
		value, ok := next()
		if !ok {
			break
		}

		fmt.Println(value)
	}
}
