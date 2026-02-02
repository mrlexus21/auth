package converter

import (
	"github.com/mrlexus21/auth/internal/model"
	modelRepo "github.com/mrlexus21/auth/internal/repository/user/model"
)

func ToUserFromRepo(user *modelRepo.User) *model.User {
	return &model.User{
		ID:          user.ID,
		Info:        ToUserInfoFromRepo(user.Info),
		Credentials: ToUserCredentialsFromRepo(user.Credentials),
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
	}
}

func ToUserInfoFromRepo(info modelRepo.UserInfo) model.UserInfo {
	return model.UserInfo{
		Name:  info.Name,
		Email: info.Email,
		Role:  info.Role,
	}
}

func ToUserCredentialsFromRepo(credentials modelRepo.UserCredentials) model.UserCredentials {
	return model.UserCredentials{
		Password:        credentials.Password,
		PasswordConfirm: credentials.PasswordConfirm,
	}
}
