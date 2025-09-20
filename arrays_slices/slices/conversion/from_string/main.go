package main

import (
	"fmt"
	"unsafe"
)

func FromString(str string) []byte {
	return unsafe.Slice(unsafe.StringData(str), len(str))
}

func main() {
	bytes := FromString("Hello, world")
	fmt.Println(bytes)
}
