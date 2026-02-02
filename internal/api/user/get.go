package user

import (
	"context"
	"log"

	"github.com/mrlexus21/auth/internal/converter"
	desc "github.com/mrlexus21/auth/pkg/user/v1"
)

func (i *Implementation) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	user, err := i.userService.Get(ctx, req.Id)
	if err != nil {
		log.Printf("Error get user: %v", err)

		return nil, err
	}

	log.Printf(
		"Get user with: id: %d, name: %s, email: %s, role: %d, created_at: %v, updated_at: %v\n",
		user.ID, user.Info.Name, user.Info.Email, user.Info.Role, user.CreatedAt, user.UpdatedAt,
	)

	return &desc.GetResponse{
		User: converter.ToUserFromService(user),
	}, nil
}
