package authservice

import (
	api "auth/pkg/auth-api"
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
)

func (s *AuthService) Register(ctx context.Context, req *api.RegisterRequest) (*api.RegisterResponse, error) {
	if err := validateRegisterRequest(req); err != nil {
		logrus.Errorf("validate register request failed: %v", err)
	}
	return &api.RegisterResponse{}, nil
}

func validateRegisterRequest(req *api.RegisterRequest) error {
	if req.GetEmail() == "" {
		return fmt.Errorf("email is required")
	}
	if req.GetPassword() == "" {
		return fmt.Errorf("password is required")
	}
	return nil
}
