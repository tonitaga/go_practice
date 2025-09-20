package main

import (
	"fmt"
	"unsafe"
)

func main() {
	bytes := []byte("Hello, world")

	str := unsafe.String(unsafe.SliceData(bytes), len(bytes))

	fmt.Println(str)

	bytes[0] = 'h'

	fmt.Println(str)
}
