package user

import (
	"context"
	"log"

	"github.com/mrlexus21/auth/internal/converter"
	desc "github.com/mrlexus21/auth/pkg/user/v1"
)

func (i *Implementation) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	userID, err := i.userService.Create(ctx, converter.ToCreateUserFromDesc(req.GetUser()))
	if err != nil {
		log.Printf("Error create user: %v", err)

		return nil, err
	}

	log.Printf("Created user with id: %d", userID)

	return &desc.CreateResponse{Id: userID}, nil
}
