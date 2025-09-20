package main

func main() {
	slice0 := make([]int, 3) // size 3 capacity 3
	slice0[5] = 10           // panic: runtime error: index out of range [5] with length 3

	slice1 := make([]int, 3, 6) // size 3 capacity 6
	slice1[5] = 10              // panic: runtime error: index out of range [5] with length 3

	// slice0[-1] = 10 // compilation error

	index := -1
	slice0[index] = 10 // panic: runtime error: index out of range [-1]

	var nilSlice []int
	if nilSlice == nil {
		nilSlice[0] = 10 // panic: runtime error: index out of range [0] with length 0
	}

	for range nilSlice {
		// ok
	}

	nilSlice = append(nilSlice, 0) // ok
}
