package main

import (
	"iter"
	"testing"
)

/*
goos: linux
goarch: amd64
cpu: AMD Ryzen 7 4800H with Radeon Graphics
BenchmarkIteratorsTransform-16          11352423               103.4 ns/op
BenchmarkSlicesTransform-16             21311937                56.49 ns/op
PASS
ok      command-line-arguments  2.385s
*/

func TransformOnIterators(data []int, action func(int) int) iter.Seq2[int, int] {
	return func(yield func(int, int) bool) {
		for i, value := range data {
			if !yield(i, action(value)) {
				break
			}
		}
	}
}

func TransformOnSlice(data []int, action func(int) int) []int {
	transformed := make([]int, len(data))
	for i, value := range data {
		transformed[i] = action(value)
	}
	return transformed
}

// go test -bench=. bench_test.go

func Action(v int) int {
	return v * v
}

func BenchmarkIteratorsTransform(b *testing.B) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}

	for b.Loop() {
		for _, value := range TransformOnIterators(data, Action) {
			_ = value
		}
	}
}

func BenchmarkSlicesTransform(b *testing.B) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	for b.Loop() {
		for _, value := range TransformOnSlice(data, Action) {
			_ = value
		}
	}
}
