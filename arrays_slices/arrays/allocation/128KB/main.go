package main

// go build -gcflags='-m'

const KB128 = 128 * (1 << 10)

func main() {
	var array1 [KB128]byte
	_ = array1

	var array2 [KB128 + 1]byte
	_ = array2
}
