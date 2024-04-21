package service

import (
	"context"
	"simplejobportal/data/request"
	"simplejobportal/helper"
	"simplejobportal/model"
	"simplejobportal/repository"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
}

func NewUserServiceImpl(userRepository repository.UserRepository) UserService {
	return &UserServiceImpl{UserRepository: userRepository}
}

// Create implements UserService.
func (b *UserServiceImpl) Create(ctx context.Context, request request.UserCreateRequest) string {

	hashPassword, _ := helper.HashPassword(request.Password)
	user := model.User{
		Username: request.Username,
		Email:    request.Email,
		Password: hashPassword,
		UserType: request.UserType,
	}
	token := b.UserRepository.Save(ctx, user)

	return token
}
