package main

// go build -gcflags='-m'

var globalArray [5]int

//go:noinline
func CreateArray() *[10]int {
	var heapArray [10]int
	return &heapArray // espace
}

func main() {
	heapArray := CreateArray()
	_ = heapArray
}
