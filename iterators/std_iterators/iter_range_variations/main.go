package main

import "fmt"

func RangeZero(yield func() bool) {
	for range 10 {
		if !yield() {
			return
		}
	}
}

func RangeOne(yield func(int) bool) {
	for value := range 10 {
		if !yield(value) {
			return
		}
	}
}

func RangeTwo(yield func(int, string) bool) {
	for value := range 10 {
		if !yield(value, fmt.Sprintf("data #%d", value)) {
			return
		}
	}
}

func main() {
	value := 0
	for range RangeZero {
		fmt.Println(value)
		value++
	}

	fmt.Println("======")

	for value := range RangeOne {
		fmt.Println(value)
	}

	fmt.Println("======")

	for value, data := range RangeTwo {
		fmt.Println(value, data)
	}
}
