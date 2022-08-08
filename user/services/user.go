package services

import (
	"context"
	"log"
	"math/rand"
	"time"

	"user/models"
)

type Server struct {
	models.UnimplementedAuthServiceServer
}

func (s *Server) GetUser(ctx context.Context, in *models.GetUserRequest) (*models.GetUserResponse, error) {
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(10)
	log.Println("Rand:", r)
	log.Println("Received:", in.Username)
	if r == 1 {
		log.Panicln("Exit")
	}
	return &models.GetUserResponse{
		Username: "Alex",
		Password: "Pass",
	}, nil
}
