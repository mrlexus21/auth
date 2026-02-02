package repository

import (
	"context"

	"github.com/mrlexus21/auth/internal/model"
)

type UserRepository interface {
	Get(ctx context.Context, id int64) (*model.User, error)
	Create(ctx context.Context, user *model.CreateUser) (int64, error)
	Update(ctx context.Context, user *model.UpdateUserInfo) error
	Delete(ctx context.Context, id int64) error
}
