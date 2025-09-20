package main

import (
	"fmt"
	"unsafe"
)

func ToString(slice []byte) string {
	if len(slice) == 0 {
		return ""
	}

	return unsafe.String(unsafe.SliceData(slice), len(slice))
}

func main() {
	str := ToString([]byte("Hello, world"))
	fmt.Println(str)
}
