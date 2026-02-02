package user

import (
	"context"
	"log"

	"github.com/mrlexus21/auth/internal/converter"
	desc "github.com/mrlexus21/auth/pkg/user/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) Update(ctx context.Context, req *desc.UpdateRequest) (*emptypb.Empty, error) {
	err := i.userService.Update(ctx, converter.ToUpdateUserInfoFromDesc(req.Info))
	if err != nil {
		log.Printf("Error update user: %v", err)

		return nil, err
	}

	log.Printf("Updated user with id: %d", req.Info.Id)

	return &emptypb.Empty{}, err
}
