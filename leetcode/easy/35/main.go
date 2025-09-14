package main

func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		middle := (left + right) / 2

		value := nums[middle]
		if value == target {
			return middle
		}

		if target > nums[middle] {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}

	return left
}

func main() {

}
