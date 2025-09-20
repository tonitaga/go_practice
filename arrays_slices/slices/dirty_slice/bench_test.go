package main

import (
	"strings"
	"testing"
	"unsafe"
)

// go test -bench=. bench_test.go

func MakeDirty(size int) []byte {
	builder := strings.Builder{}
	builder.Grow(size)

	pointer := unsafe.StringData(builder.String())
	return unsafe.Slice(pointer, size)
}

func BenchmarkMakeUsual(b *testing.B) {
	var result []byte

	for range b.N {
		result = make([]byte, 1<<20)
	}

	_ = result
}

func BenchmarkMakeDirty(b *testing.B) {
	var result []byte

	for range b.N {
		result = MakeDirty(1 << 20)
	}

	_ = result
}
