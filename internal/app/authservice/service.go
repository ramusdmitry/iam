package authservice

import api "auth/pkg/auth-api"

type AuthService struct {
	api.UnimplementedAuthServiceServer
}

func NewAuthService() *AuthService {
	return &AuthService{}
}
