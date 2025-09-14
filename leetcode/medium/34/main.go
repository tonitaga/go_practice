package main

func searchRange1(nums []int, target int) []int {
	left, right := 0, len(nums)-1

	for left <= right {
		middle := (left + right) / 2

		if nums[middle] == target {
			lhs, rhs := middle, middle
			for ; lhs >= 0; lhs-- {
				if nums[lhs] != target {
					break
				}
			}

			for ; rhs < len(nums); rhs++ {
				if nums[rhs] != target {
					break
				}
			}

			return []int{lhs + 1, rhs - 1}
		}

		if nums[middle] < target {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}

	return []int{-1, -1}
}

func binarySearch(nums []int, target int, go_left bool) int {
	left, right := 0, len(nums)-1

	index := -1
	for left <= right {
		middle := (left + right) / 2

		if nums[middle] == target {
			index = middle
			if go_left {
				right = middle - 1
			} else {
				left = middle + 1
			}
		} else if nums[middle] < target {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}

	return index
}

func searchRange2(nums []int, target int) []int {
	return []int{
		binarySearch(nums, target /*go_left=*/, true),
		binarySearch(nums, target /*go_left=*/, false),
	}
}

func main() {
}
