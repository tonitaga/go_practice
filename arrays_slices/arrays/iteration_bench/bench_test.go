package main

import "testing"

const MB64 = 64 * (1 << 20)

// go test -bench=. bench_test.go

func BenchmarkIterationArrayUsual(b *testing.B) {
	var array [MB64]byte
	for range b.N {
		for range array {
		}
	}
}

func BenchmarkIterationArrayBySlice(b *testing.B) {
	var array [MB64]byte
	for range b.N {
		for range &array {
		}
	}
}

func BenchmarkIterationArrayByPointer(b *testing.B) {
	var array [MB64]byte
	for range b.N {
		for range array[:] {
		}
	}
}
