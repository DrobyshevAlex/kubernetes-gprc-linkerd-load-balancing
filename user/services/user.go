package services

import (
	"context"
	"log"
	"math/rand"
	"time"

	"user/models"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	models.UnimplementedAuthServiceServer
	V int
}

func (s *Server) GetUser(ctx context.Context, in *models.GetUserRequest) (*models.GetUserResponse, error) {
	log.Println("")
	log.Println("***** GetUser ******")
	log.Println("Received:", in.Username)
	if in.Username == "crash" {
		rand.Seed(time.Now().UnixNano())
		r := rand.Intn(10)
		log.Println("Rand:", r)
		if r == 1 {
			log.Println("***** GetUser crashed ******")
			log.Panicln("Exit")
		}
	} else {
		s.V++
		if s.V%2 == 0 {
			log.Println("***** GetUser status Internal error ******")
			return nil, status.Error(codes.Internal, in.Username)
		}
	}
	log.Println("***** GetUser status OK ******")
	return &models.GetUserResponse{
		Username: "Alex",
		Password: "Pass",
	}, nil
}
