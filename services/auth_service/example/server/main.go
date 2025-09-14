package main

import (
	"fmt"

	"github.com/go_practice/sandbox/auth_service/internal/transport/grpc/server"
)

func main() {
	server := server.NewDefault()
	if err := server.Listen(":8088"); err != nil {
		fmt.Println(err)
	}
}
