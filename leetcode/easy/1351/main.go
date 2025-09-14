package main

func countNegatives(grid [][]int) int {
	negativesCount := 0
	for _, array := range grid {
		size := len(array)
		left, right := 0, size-1

		for left <= right {
			middle := (left + right) / 2

			if array[middle] < 0 {
				right = middle - 1
			} else {
				left = middle + 1
			}
		}

		negativesCount += (size - left)
	}

	return negativesCount
}

func main() {

}
