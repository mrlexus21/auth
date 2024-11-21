package main

import (
	"context"
	"flag"
	"github.com/fatih/color"
	"github.com/mrlexus21/auth/internal/config"
	user_v1 "github.com/mrlexus21/auth/pkg/user/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", ".env", "path to config file")
}

const userID int64 = 1

func main() {
	flag.Parse()

	// Считываем переменные окружения
	if err := config.Load(configPath); err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	grpcConfig, err := config.NewGRPCConfig()
	if err != nil {
		log.Fatalf("failed to get grpc config: %v", err)
	}

	conn, err := grpc.NewClient(grpcConfig.Address(), grpc.WithTransportCredentials(insecure.NewCredentials()))
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

	log.Printf("User info:\n%s", color.GreenString("%+v", r.String()))
}
