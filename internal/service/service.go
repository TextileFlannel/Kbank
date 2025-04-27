package service

import (
	"kbank/api"
	"kbank/internal/storage"
)

type AuthService struct {
	api.UnimplementedAuthServiceServer
	storage storage.Storage
}
