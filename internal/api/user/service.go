package user

import (
	"github.com/mrlexus21/auth/internal/service"
	desc "github.com/mrlexus21/auth/pkg/user/v1"
)

type Implementation struct {
	desc.UnimplementedUserV1Server
	userService service.UserService
}

func NewImplementation(userService service.UserService) *Implementation {
	return &Implementation{
		userService: userService,
	}
}
