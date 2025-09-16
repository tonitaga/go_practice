package main

func main() {
	someMap := map[int]int{1: 1}
	for range someMap {
		// Handle
	}

	someSlice := []int{1, 2, 3}
	for range someSlice {
		// Handle
	}

	someArray := [...]int{1, 2, 3}
	for range someArray {
		// Handle
	}

	someString := "Hello"
	for range someString {
		// Handle
	}

	someChan := make(chan int)

	defer close(someChan)
	for range someChan {
		// Handle
	}

	someDigit := 10
	for range someDigit {
		// Handle
	}
}
