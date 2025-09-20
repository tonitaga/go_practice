package main

import "fmt"

func main() {
	slice0 := make([]int, 5, 10)

	for i := range slice0 {
		slice0[i] = i
	}

	fmt.Println(slice0)
	fmt.Println("size:", len(slice0), "capacity:", cap(slice0))

	for i := range slice0 {
		slice := slice0[i+1:]

		fmt.Println(slice)
		fmt.Println("size:", len(slice), "capacity:", cap(slice))
	}
}
