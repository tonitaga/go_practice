package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go_practice/sandbox/auth_service/pkg/grpc/client"
	auth "github.com/go_practice/sandbox/auth_service/pkg/grpc/gen"
)

func main() {
	client, err := client.NewDefault(":8088")
	if err != nil {
		panic(err)
	}

	context, cancel := context.WithTimeout(context.Background(), time.Second)

	defer func() {
		cancel()
		client.Close()
	}()

	response, err := client.Get().Login(context, &auth.LoginRequest{
		Username: "tonitaga",
		Password: "1234",
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(response)
}
