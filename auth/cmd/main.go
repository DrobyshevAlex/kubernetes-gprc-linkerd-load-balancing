package main

import (
	"context"
	"flag"
	"log"
	"strconv"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"auth/models"
)

var (
	addr = flag.String("addr", "app-user:50051", "the address to connect to")
	//addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := models.NewAuthServiceClient(conn)

	for i := 0; ; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		r, err := c.GetUser(ctx, &models.GetUserRequest{Username: "User " + strconv.Itoa(i)})
		if err != nil {
			log.Printf("not exec: %v", err)
		} else {
			log.Printf("Success: %s", r.Password)
		}
		time.Sleep(time.Second * 1)
	}
}
