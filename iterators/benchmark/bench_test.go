package main

import "testing"

const IterationsInOneTest = 10

func Range(yield func(int) bool) {
	for value := range IterationsInOneTest {
		if !yield(value) {
			break
		}
	}
}

func IteratorRange() {
	for value := range Range {
		_ = value
	}
}

func UsualRange() {
	for value := range IterationsInOneTest {
		_ = value
	}
}

// go test -bench=. bench_test.go

func BenchmarkIteratorRange(b *testing.B) {
	for b.Loop() {
		IteratorRange()
	}
}

func BenchmarkUsualRange(b *testing.B) {
	for b.Loop() {
		UsualRange()
	}
}
