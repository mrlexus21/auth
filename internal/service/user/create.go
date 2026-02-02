package user

import (
	"context"

	"github.com/mrlexus21/auth/internal/model"
)

func (s *serv) Create(ctx context.Context, user *model.CreateUser) (int64, error) {
	userID, err := s.userRepository.Create(ctx, user)
	if err != nil {
		return 0, err
	}

	return userID, nil
}
