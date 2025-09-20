package main

const KB64 = 64 * (1 << 10)

// go build -gcflags='-m' .

func main() {
	slice0 := make([]byte, KB64)
	slice1 := make([]byte, KB64+1)

	_, _ = slice0, slice1
}
