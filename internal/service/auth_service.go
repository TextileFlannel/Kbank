package service

import (
	"context"
	"kbank/api"

	"golang.org/x/crypto/bcrypt"
	"kbank/internal/storage"
)

func NewAuthService(storage storage.Storage) *AuthService {
	return &AuthService{
		storage: storage,
	}
}

func (s *AuthService) Register(ctx context.Context, req *api.RegisterRequest) (*api.RegisterResponse, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return &api.RegisterResponse{
			Success: false,
			Error:   "internal error",
		}, nil
	}

	err = s.storage.CreateUser(req.Login, hashedPassword)
	if err != nil {
		return &api.RegisterResponse{
			Success: false,
			Error:   err.Error(),
		}, nil
	}

	return &api.RegisterResponse{Success: true}, nil
}

func (s *AuthService) Login(ctx context.Context, req *api.LoginRequest) (*api.LoginResponse, error) {
	user, err := s.storage.GetUser(req.Login)
	if err != nil {
		return &api.LoginResponse{
			Error: "invalid credentials",
		}, nil
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(req.Password)); err != nil {
		return &api.LoginResponse{
			Error: "invalid credentials",
		}, nil
	}

	return &api.LoginResponse{
		Token: "generated-jwt-token",
	}, nil
}
