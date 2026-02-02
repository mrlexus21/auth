package converter

import (
	"github.com/mrlexus21/auth/internal/model"
	desc "github.com/mrlexus21/auth/pkg/user/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToUserFromService(user *model.User) *desc.User {
	var updatedAt *timestamppb.Timestamp
	if user.UpdatedAt.Valid {
		updatedAt = timestamppb.New(user.UpdatedAt.Time)
	}

	return &desc.User{
		Id:          user.ID,
		Info:        ToUserInfoFromService(user.Info),
		Credentials: ToUserCredentialsFromService(user.Credentials),
		CreatedAt:   timestamppb.New(user.CreatedAt),
		UpdatedAt:   updatedAt,
	}
}

func ToUserInfoFromService(info model.UserInfo) *desc.UserInfo {
	return &desc.UserInfo{
		Name:  info.Name,
		Email: info.Email,
		Role:  desc.Roles(info.Role),
	}
}

func ToUserCredentialsFromService(credentials model.UserCredentials) *desc.UserCredentials {
	return &desc.UserCredentials{
		Password:        credentials.Password,
		PasswordConfirm: credentials.PasswordConfirm,
	}
}

func ToCreateUserFromDesc(user *desc.CreateUser) *model.CreateUser {
	return &model.CreateUser{
		Info:        ToUserInfoFromDesc(user.Info),
		Credentials: ToUserCredentialsFromDesc(user.Credentials),
	}
}

func ToUserInfoFromDesc(info *desc.UserInfo) model.UserInfo {
	return model.UserInfo{
		Name:  info.Name,
		Email: info.Email,
		Role:  int32(info.Role),
	}
}

func ToUserCredentialsFromDesc(credentials *desc.UserCredentials) model.UserCredentials {
	return model.UserCredentials{
		Password:        credentials.Password,
		PasswordConfirm: credentials.PasswordConfirm,
	}
}

func ToUpdateUserInfoFromDesc(info *desc.UpdateUserInfo) *model.UpdateUserInfo {
	var name, email string

	if info.Name != nil {
		name = info.Name.Value
	}
	if info.Email != nil {
		email = info.Email.Value
	}

	return &model.UpdateUserInfo{
		ID:    info.Id,
		Name:  &name,
		Email: &email,
	}
}
