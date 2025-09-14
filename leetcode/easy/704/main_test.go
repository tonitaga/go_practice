package main

import "testing"

func TestSearchExists(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8}

	for i, num := range nums {
		index := search(nums, num)

		if index != i {
			t.Errorf("Expected index '%d' on number: '%d'. Got '%d'", index, num, i)
		}
	}
}

func TestSearchNotExists(t *testing.T) {
	nums := []int{1, 3, 5, 7, 9}

	for _, num := range nums {
		index := search(nums, num+1)

		if index != -1 {
			t.Errorf("Expected index '-1' on number: '%d'. Got '%d'", num+1, index)
		}
	}
}
