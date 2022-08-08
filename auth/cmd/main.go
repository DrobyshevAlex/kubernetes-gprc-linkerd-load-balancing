package main

import (
	"context"
	"log"
	"time"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"github.com/gin-gonic/gin"

	"auth/models"
)

func main() {
	addrGrpc := os.Getenv("GRPC_ADDR")
	if len(addrGrpc) == 0 {
		addrGrpc = "localhost:50051"
	}
	
	conn, err := grpc.Dial(addrGrpc, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	connAuth := models.NewAuthServiceClient(conn)

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		r, err := connAuth.GetUser(ctx, &models.GetUserRequest{Username: "User"})
		if err != nil {
			log.Printf("not exec: %v", err)
			c.AbortWithStatusJSON(503, gin.H{"error": err})
		} else {
			log.Printf("Success: %s", r.Password)
			c.JSON(200, gin.H{
				"message": r.Password,
			})
		}
	})

	addr := os.Getenv("AUTH_ADDR")
	if len(addr) == 0 {
		addr = "0.0.0.0:8080"
	}
	r.Run(addr)
}
