package main

func main() {
	slice0 := []int{1, 2}
	slice1 := [][]int{{1, 2, 3}, {1, 2, 3}}

	_, _ = slice0, slice1

	slice2 := []int{4: 5}       // [0,0,0,0,5] size 5 capacity 5
	slice3 := make([]int, 4)    // [0,0,0,0]   size 4 capacity 4
	slice4 := make([]int, 4, 5) // [0,0,0,0,0] size 4 capacity 5

	_, _, _ = slice2, slice3, slice4
}
