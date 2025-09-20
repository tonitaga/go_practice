package main

import (
	"fmt"
)

func main() {
	slice0 := make([]int, 10)
	for i := range slice0 {
		slice0[i] = i
	}

	var slice2 []int = nil

	slice2 = append(slice2, slice0...)
	fmt.Println(slice2)

	// Удаление серединного элемента среза. Например с индексом 2

	slice2 = append(slice2[:2], slice2[3:]...)
	fmt.Println(slice2)
}
