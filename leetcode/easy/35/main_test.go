package main

import (
	"testing"
)

func TestSearchInsertExists(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8}

	for i, num := range nums {
		index := searchInsert(nums, num)

		if index != i {
			t.Errorf("Expected index '%d' on number: '%d'. Got '%d'", index, num, i)
		}
	}
}

func TestSearchInsertNotExists(t *testing.T) {
	nums := []int{1, 3, 5, 7, 9}
	targets := []int{0, 2, 4, 6, 8, 10}
	results := []int{0, 1, 2, 3, 4, 5}

	for i, num := range targets {
		index := searchInsert(nums, num)

		if index != results[i] {
			t.Errorf("Expected index '%d' on number: '%d'. Got '%d'", results[i], num, index)
		}
	}
}
