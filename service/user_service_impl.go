package service

import (
	"context"
	"errors"
	"simplejobportal/auth"
	"simplejobportal/data/request"
	"simplejobportal/helper"
	"simplejobportal/model"
	"simplejobportal/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
}

func NewUserServiceImpl(userRepository repository.UserRepository) UserService {
	return &UserServiceImpl{UserRepository: userRepository}
}

// Create implements UserService.
func (b *UserServiceImpl) Create(ctx context.Context, request request.UserCreateRequest) (string, error) {

	hashPassword, _ := helper.HashPassword(request.Password)
	user := model.User{
		Username: request.Username,
		Email:    request.Email,
		Password: hashPassword,
		UserType: request.UserType,
	}
	token, err := b.UserRepository.Save(ctx, user)

	return token, err
}

// Login implements UserService.
func (b *UserServiceImpl) Login(ctx context.Context, request request.UserLoginRequest) (string, error) {

	user, err := b.UserRepository.GetByEmail(ctx, request.Email)
	helper.PanicIfError(err)
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		return "", errors.New("wrong password")
	}
	token, err := auth.CreateJWT(user.Id, user.UserType)
	helper.PanicIfError(err)
	return token, nil
}
