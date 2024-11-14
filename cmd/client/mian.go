package main

import (
	"context"
	"github.com/fatih/color"
	user_v1 "github.com/mrlexus21/auth/pkg/user/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

const (
	address = "localhost:50051"
	userID  = 1
)

func main() {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalf("Failed to close connection: %v", err)
		}
	}(conn)

	c := user_v1.NewUserV1Client(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.Get(ctx, &user_v1.GetRequest{Id: userID})
	if err != nil {
		log.Fatalf("Failed to get user by id: %v", err)
	}

	log.Printf(color.RedString("User info:\n"), color.GreenString("%+v", r.String()))
}
