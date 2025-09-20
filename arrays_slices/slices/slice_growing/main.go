package main

import (
	"fmt"
)

func main() {
	slice := make([]int, 1)

	lastCapacity := cap(slice)

	for range 1 << 13 {
		slice = append(slice, 0)

		currentCapacity := cap(slice)
		if currentCapacity != lastCapacity {
			fmt.Println("Capacity:", currentCapacity, "Size:", len(slice), "Factor:", float64(currentCapacity)/float64(lastCapacity))
			lastCapacity = currentCapacity
		}
	}
}
