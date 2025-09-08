package main

import "testing"

func TestIsValidCorrectCases(t *testing.T) {
	correctCases := []string{
		"", "[]", "{}", "[]", "{}()[]", "{([])}", "{}([])", "[()]{}", "[{}()]",
	}

	for _, testCase := range correctCases {
		if !isValid(testCase) {
			t.Errorf("Expected 'true' on testcase: '%s'\n", testCase)
		}
	}
}

func TestIsValidFailCases(t *testing.T) {
	correctCases := []string{
		"{", "}", "{}(){", "[{]}",
	}

	for _, testCase := range correctCases {
		if isValid(testCase) {
			t.Errorf("Expected 'false' on testcase: '%s'\n", testCase)
		}
	}
}
