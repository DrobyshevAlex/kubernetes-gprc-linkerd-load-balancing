package services

import (
	"context"
	"log"

	"user/models"
)

type Server struct {
	models.UnimplementedAuthServiceServer
	V int
}

func (s *Server) GetUser(ctx context.Context, in *models.GetUserRequest) (*models.GetUserResponse, error) {
	s.V++

	log.Println("V:", s.V)
	log.Println("Received:", in.Username)
	if s.V == 2 {
		log.Panicln("Exit")
	}
	return &models.GetUserResponse{
		Username: "Alex",
		Password: "Pass",
	}, nil
}
