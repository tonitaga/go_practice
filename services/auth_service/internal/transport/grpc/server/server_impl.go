package server

import (
	"context"

	auth "github.com/go_practice/sandbox/auth_service/pkg/grpc/gen"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type authServiceServerImpl struct {
	auth.UnimplementedAuthServiceServer
}

func (s *authServiceServerImpl) Register(ctx context.Context, r *auth.RegisterRequest) (*auth.RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}

func (s *authServiceServerImpl) Login(ctx context.Context, r *auth.LoginRequest) (*auth.LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}

func (s *authServiceServerImpl) Logout(ctx context.Context, r *auth.LogoutRequest) (*auth.LogoutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logout not implemented")
}

func (s *authServiceServerImpl) ValidateToken(ctx context.Context, r *auth.ValidateTokenRequest) (*auth.ValidateTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateToken not implemented")
}

func (s *authServiceServerImpl) RefreshToken(ctx context.Context, r *auth.RefreshTokenRequest) (*auth.RefreshTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshToken not implemented")
}
