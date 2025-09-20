package main

func main() {
	var arr0 [5]int    // [0,0,0,0,0]
	var arr1 [2][3]int // [[0,0,0],[0,0,0]]

	_, _ = arr0, arr1

	arr2 := [...]int{1, 2, 3}     // [1,2,3]
	arr3 := [5]int{1, 2, 3}       // [1,2,3,0,0]
	arr4 := [5]int{4: 5}          // [0,0,0,0,5]
	arr5 := [5]int{2: 5, 6, 1: 9} // [0,9,5,6,0]

	_, _, _, _ = arr2, arr3, arr4, arr5
}
