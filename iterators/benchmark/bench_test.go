package main

import "testing"

/*
goos: linux
goarch: amd64
cpu: AMD Ryzen 7 4800H with Radeon Graphics
BenchmarkIteratorRange-16       313387570                3.849 ns/op
BenchmarkUsualRange-16          310089063                3.869 ns/op
PASS
ok      command-line-arguments  2.412s
*/

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
