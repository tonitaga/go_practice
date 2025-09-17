package main

import (
	"fmt"
	"iter"
)

func Fibonacci(n int) iter.Seq[int64] {
	return func(yield func(int64) bool) {
		count := 0

		var lhs int64 = 0
		var rhs int64 = 1

		for count <= n {
			if !(yield(lhs)) {
				return
			}

			lhs, rhs = rhs, lhs+rhs
			count++
		}
	}
}

func main() {
	for value := range Fibonacci(50) {
		fmt.Println(value)
	}
}
