package main

import (
	"fmt"
	"iter"
)

func CatchPanic() {
	if err := recover(); err != nil {
		fmt.Println("Caugth panic:", err)
	}
}

func Range(size int) iter.Seq[int] {
	defer CatchPanic()

	return func(yield func(int) bool) {
		for value := range size {
			if !yield(value) {
				return
			}
		}
	}
}

func main() {
	defer CatchPanic() // Можно ловить здесь

	for value := range Range(10) {
		if value == 5 {
			panic("panic")
		}
		fmt.Println(value)
	}
}
