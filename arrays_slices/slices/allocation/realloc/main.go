package main

// go build -gcflags='-m' .

func main() {
	slice := make([]byte, 0, 3)
	println("slice:", slice, "header address:", &slice)

	slice = append(slice, 1, 1, 1)
	println("slice:", slice, "header address:", &slice)

	slice = append(slice, 1)
	println("slice:", slice, "header address:", &slice)
}
