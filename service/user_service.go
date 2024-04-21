package service

import (
	"context"
	"simplejobportal/data/request"
)

type UserService interface {
	Create(ctx context.Context, request request.UserCreateRequest) string
}
