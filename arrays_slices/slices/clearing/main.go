package main

import "fmt"

func main() {
	slice0 := make([]int, 0, 10)
	slice0 = append(slice0, 1, 2, 3, 4, 5)

	slice0 = nil
	fmt.Println("size:", len(slice0), "cap:", cap(slice0))

	_ = slice0

	slice1 := make([]int, 0, 10)

	slice1 = slice1[:0]
	fmt.Println("size:", len(slice1), "cap:", cap(slice1))

	_ = slice1

	slice2 := make([]int, 0, 10)
	slice2 = append(slice2, 1, 2, 3, 4, 5)

	clear(slice2)
	fmt.Println(slice2)
	fmt.Println("size:", len(slice2), "cap:", cap(slice2))

	slice2 = slice2[:0]
	slice2 = append(slice2, 1, 2, 3, 4, 5)

	clear(slice2[1:3])
	fmt.Println(slice2)
}
