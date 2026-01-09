// Package main реализует серверную часть приложения аутентификации.
// Обеспечивает обработку gRPC-запросов, управление подключением к базе данных и бизнес-логику аутентификации.
package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/mrlexus21/auth/internal/config"
	user_v1 "github.com/mrlexus21/auth/pkg/user/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", ".env", "path to config file")
}

type server struct {
	user_v1.UnimplementedUserV1Server
	pool *pgxpool.Pool
}

func main() {
	flag.Parse()
	ctx := context.Background()

	// Считываем переменные окружения
	if err := config.Load(configPath); err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	grpcConfig, err := config.NewGRPCConfig()
	if err != nil {
		log.Fatalf("failed to get grpc config: %v", err)
	}

	pgConfig, err := config.NewPGConfig()
	if err != nil {
		log.Fatalf("failed to get pg config: %v", err)
	}

	lis, err := net.Listen("tcp", grpcConfig.Address())
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Создаем пул соединений с базой данных
	pool, err := pgxpool.Connect(ctx, pgConfig.DSN())
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer pool.Close()

	s := grpc.NewServer()
	reflection.Register(s)
	user_v1.RegisterUserV1Server(s, &server{pool: pool})

	log.Printf("Server listening at %s", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func (s *server) Get(ctx context.Context, req *user_v1.GetRequest) (*user_v1.GetResponse, error) {
	query, args, err := sq.Select("id", "name", "email", "role", "created_at", "updated_at").
		From("users").
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id": req.GetId()}).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build query: %v", err)
	}

	var id int64
	var name, email string
	var role int32
	var createdAt time.Time
	var updatedAt sql.NullTime

	if err := s.pool.QueryRow(ctx, query, args...).Scan(&id, &name, &email, &role, &createdAt, &updatedAt); err != nil {
		return nil, fmt.Errorf("failed to select user: %v", err)
	}

	log.Printf("id: %d, name: %s, email: %s, role: %d, created_at: %v, updated_at: %v\n", id, name, email, role, createdAt, updatedAt)

	var updatedAtTime *timestamppb.Timestamp
	if updatedAt.Valid {
		updatedAtTime = timestamppb.New(updatedAt.Time)
	}

	return &user_v1.GetResponse{
		Id:        id,
		Name:      name,
		Email:     email,
		Role:      user_v1.Roles(role),
		CreatedAt: timestamppb.New(createdAt),
		UpdatedAt: updatedAtTime,
	}, nil
}

func (s *server) Create(ctx context.Context, req *user_v1.CreateRequest) (*user_v1.CreateResponse, error) {
	query, args, err := sq.Insert("users").
		PlaceholderFormat(sq.Dollar).
		Columns("name", "email", "password", "password_confirm", "role").
		Values(req.GetName(), req.GetEmail(), req.GetPassword(), req.GetPasswordConfirm(), req.GetRole()).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build query: %v", err)
	}

	var userID int64
	if err := s.pool.QueryRow(ctx, query, args...).Scan(&userID); err != nil {
		return nil, fmt.Errorf("failed to insert user: %v", err)
	}

	log.Printf("inserted user with id: %d", userID)

	return &user_v1.CreateResponse{Id: userID}, nil
}

func (s *server) Update(ctx context.Context, req *user_v1.UpdateRequest) (*emptypb.Empty, error) {
	builderUpdate := sq.Update("users").
		PlaceholderFormat(sq.Dollar).
		Set("updated_at", time.Now())

	if name := req.GetName(); name != nil && name.GetValue() != "" {
		builderUpdate = builderUpdate.Set("name", name.GetValue())
	}
	if email := req.GetEmail(); email != nil && email.GetValue() != "" {
		builderUpdate = builderUpdate.Set("email", email.GetValue())
	}

	builderUpdate = builderUpdate.Where(sq.Eq{"id": req.GetId()})

	query, args, err := builderUpdate.ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build query: %v", err)
	}

	res, err := s.pool.Exec(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to update user: %v", err)
	}

	log.Printf("updated %d rows", res.RowsAffected())

	return &emptypb.Empty{}, nil
}

func (s *server) Delete(ctx context.Context, req *user_v1.DeleteRequest) (*emptypb.Empty, error) {
	query, args, err := sq.Delete("users").
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id": req.GetId()}).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build query: %v", err)
	}

	res, err := s.pool.Exec(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to delete user: %v", err)
	}

	log.Printf("deleted %d rows", res.RowsAffected())

	return &emptypb.Empty{}, nil
}
