package main

import "testing"

func TestNextGreatestLetter(t *testing.T) {
	testCases := []string{
		"aabbbccc",
		"cfj",
		"cfj",
		"xxyyy",
	}

	targets := []byte{
		'd',
		'a',
		'c',
		'z',
	}

	results := []byte{
		'a', // Not found. Returns letters[0]
		'c',
		'f',
		'x',
	}

	for i, testCase := range testCases {
		result := nextGreatestLetter([]byte(testCase), targets[i])
		if result != results[i] {
			t.Errorf("Expected result '%d' on target: '%d'. Got '%d'", results[i], targets[i], result)
		}
	}
}
