package client

import (
	auth "github.com/go_practice/sandbox/auth_service/pkg/grpc/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type AuthServiceClient struct {
	connection *grpc.ClientConn
	client     auth.AuthServiceClient
}

func NewDefault(address string) (*AuthServiceClient, error) {
	connection, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	authServiceClient := &AuthServiceClient{
		connection: connection,
		client:     auth.NewAuthServiceClient(connection),
	}

	return authServiceClient, nil
}

func (c *AuthServiceClient) Get() auth.AuthServiceClient {
	return c.client
}

func (c *AuthServiceClient) Close() {
	c.connection.Close()
}
