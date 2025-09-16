package main

import "fmt"

func BadRange(yield func(int) bool) {
	for value := range 10 {
		yield(value) // yield may be called again after returning false
	}
}

func main() {
	for value := range BadRange {
		fmt.Println(value)

		if value == 5 {
			break
		}
	}
}
