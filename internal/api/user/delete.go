package user

import (
	"context"
	"log"

	desc "github.com/mrlexus21/auth/pkg/user/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	err := i.userService.Delete(ctx, req.Id)
	if err != nil {
		log.Printf("Error delete user: %v", err)

		return nil, err
	}

	log.Printf("Deleted user with id: %d", req.Id)

	return &emptypb.Empty{}, err
}
