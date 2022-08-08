package main

import (
	"context"
	"flag"
	"log"
	"time"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"github.com/gin-gonic/gin"

	"auth/models"
)

var (
	//addr = flag.String("addr", "app-user:50051", "the address to connect to")
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
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
