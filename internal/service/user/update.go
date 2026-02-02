package user

import (
	"context"

	"github.com/mrlexus21/auth/internal/model"
)

func (s *serv) Update(ctx context.Context, info *model.UpdateUserInfo) error {
	err := s.userRepository.Update(ctx, info)

	return err
}
