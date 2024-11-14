package main

import (
	"context"
	"crypto/rand"
	"fmt"
	"github.com/brianvoe/gofakeit"
	user_v1 "github.com/mrlexus21/auth/pkg/user/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"math/big"
	"net"
)

const grpcPort = 50051

type server struct {
	user_v1.UnimplementedUserV1Server
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	user_v1.RegisterUserV1Server(s, &server{})

	log.Printf("Server listening at %s", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func (s *server) Get(_ context.Context, req *user_v1.GetRequest) (*user_v1.GetResponse, error) {
	log.Printf("Get user: %d", req.GetId())

	return &user_v1.GetResponse{
		Id:        req.GetId(),
		Name:      gofakeit.Name(),
		Email:     gofakeit.Email(),
		Role:      user_v1.Roles_ADMIN,
		CreatedAt: timestamppb.New(gofakeit.Date()),
		UpdatedAt: timestamppb.New(gofakeit.Date()),
	}, nil
}

func (s *server) Create(_ context.Context, req *user_v1.CreateRequest) (*user_v1.CreateResponse, error) {
	log.Printf("Create user: %v", req.String())

	id, err := rand.Int(rand.Reader, big.NewInt(1<<63-1))
	if err != nil {
		log.Fatalf("Failed user id generate: %v", err)
		return &user_v1.CreateResponse{
			Id: 0,
		}, nil
	}
	return &user_v1.CreateResponse{Id: id.Int64()}, nil
}

func (s *server) Update(_ context.Context, req *user_v1.UpdateRequest) (*emptypb.Empty, error) {
	log.Printf("Update user: %v", req.String())

	return &emptypb.Empty{}, nil
}

func (s *server) Delete(_ context.Context, req *user_v1.DeleteRequest) (*emptypb.Empty, error) {
	log.Printf("Delete user: %d", req.GetId())

	return &emptypb.Empty{}, nil
}
