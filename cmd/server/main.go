package main

import (
	"log"
	"net"

	"kbank/api"
	"kbank/internal/service"
	"kbank/internal/storage"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	store := storage.NewInMemStorage()

	srv := service.NewAuthService(store)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	api.RegisterAuthServiceServer(s, srv)
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
