package repository

import (
	"context"
	"simplejobportal/model"
)

type UserRepository interface {
	Save(ctx context.Context, user model.User) (string, error)
	GetByEmail(ctx context.Context, email string) (model.User, error)
}
