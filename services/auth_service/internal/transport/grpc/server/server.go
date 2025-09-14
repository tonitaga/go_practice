package server

import (
	"fmt"
	"log"
	"net"

	auth "github.com/go_practice/sandbox/auth_service/pkg/grpc/gen"
	"google.golang.org/grpc"
)

type AuthServiceServer struct {
	server  *grpc.Server
	service authServiceServerImpl
}

func NewDefault() *AuthServiceServer {
	authServer := &AuthServiceServer{
		server:  grpc.NewServer(),
		service: authServiceServerImpl{},
	}

	auth.RegisterAuthServiceServer(authServer.server, &authServer.service)
	return authServer
}

func (s *AuthServiceServer) Listen(address string) error {
	op := "AuthServiceServer.Listen"

	listener, err := net.Listen("tcp", address)
	if err != nil {
		return fmt.Errorf("%s: %v", op, err)
	}

	log.Printf("%s: Listening %s", op, address)

	if err := s.server.Serve(listener); err != nil {
		return fmt.Errorf("%s: Failed to serve. Cause: %v", op, err)
	}

	return nil
}
