package repository

import (
	"context"
	"simplejobportal/model"
)

type UserRepository interface {
	Save(ctx context.Context, user model.User) string
}
