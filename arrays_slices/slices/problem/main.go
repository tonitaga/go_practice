package main

import "fmt"

func proc1(data []int) {
	if len(data) > 2 {
		data[2] = 5
	}
}

func proc2(data []int) {
	data = append(data, 6) // this value of data is never used (SA4006)go-staticcheck
}

func main() {
	data := make([]int, 0, 10)
	data = append(data, 1, 2, 3, 4, 5)

	fmt.Println(data)

	proc1(data)

	fmt.Println(data) // [1,2,5,4,5]

	proc2(data)

	fmt.Println(data) // [1,2,5,4,5] // PROBLEM

	data = data[:6]
	fmt.Println(data) // [1,2,5,4,5,6]
}
