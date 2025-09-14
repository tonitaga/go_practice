package main

/*
t = 8

5,7,7,8,8,8,8,8,8,8,8,10
L         M            R
L   M1    M     M2     R

	L     M     M      R

r = [3,10]
*/
func searchRange(nums []int, target int) []int {
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

func main() {
}
