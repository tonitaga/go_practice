package main

import "testing"

func TestSearchRange(t *testing.T) {
	nums := []int{5, 7, 7, 8, 8, 10}
	targets := []int{8, 6, 10}
	results := [][]int{
		{3, 4},
		{-1, -1},
		{5, 5},
	}

	for i, target := range targets {
		result := searchRange1(nums, target)
		if len(result) != 2 {
			t.Errorf("Result of searchRange must be array of two elements. Got: %d elements", len(result))
			continue
		}

		for j, value := range result {
			if value != results[i][j] {
				t.Errorf("Wrong interval: Got: [%d, %d]. Expected: [%d, %d]", result[0], result[1], results[i][0], results[i][1])
				continue
			}
		}
	}

	for i, target := range targets {
		result := searchRange2(nums, target)
		if len(result) != 2 {
			t.Errorf("Result of searchRange must be array of two elements. Got: %d elements", len(result))
			continue
		}

		for j, value := range result {
			if value != results[i][j] {
				t.Errorf("Wrong interval: Got: [%d, %d]. Expected: [%d, %d]", result[0], result[1], results[i][0], results[i][1])
				continue
			}
		}
	}
}
