package main

import "testing"

func TestCountNegatives(t *testing.T) {
	testCase := [][]int{
		{4, 3, 2, -1},
		{3, 2, 1, -1},
		{1, 1, -1, -2},
		{-1, -1, -2, -3},
	}

	expectedResult := 8

	result := countNegatives(testCase)
	if result != expectedResult {
		t.Errorf("Extected '%d' result on testcase: %v. Got: %d", expectedResult, testCase, result)
	}
}
