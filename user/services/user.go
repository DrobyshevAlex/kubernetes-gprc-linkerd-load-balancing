package services

import (
	"context"
	"log"

	"user/models"
)

type Server struct {
	models.UnimplementedAuthServiceServer
}

func (s *Server) GetUser(ctx context.Context, in *models.GetUserRequest) (*models.GetUserResponse, error) {
	log.Println("Received:", in.Username)
	return &models.GetUserResponse{
		Username: "Alex",
		Password: "Pass",
	}, nil
}
