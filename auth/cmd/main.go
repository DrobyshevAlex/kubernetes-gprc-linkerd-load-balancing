package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

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
	r.GET("/test", func(c *gin.Context) {
		resp, err := http.Get("http://app-user")
		fmt.Println(resp)
		fmt.Println(err)
		c.JSON(200, gin.H{
			"message": "OK",
		})
	})
	r.GET("/", func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		//outgoingContext := metadata.AppendToOutgoingContext(ctx, "x-envoy-retry-grpc-on", "5xx,cancelled,deadline-exceeded,internal,resource-exhausted,unavailable", "x-envoy-max-retries", "3", "test-key", "test-val", "x-envoy-upstream-rq-timeout-ms", "15000")
		defer cancel()
		fmt.Println("QUERY GRPC")
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
