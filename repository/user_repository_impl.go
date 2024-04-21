package repository

import (
	"context"
	"database/sql"
	"simplejobportal/auth"
	"simplejobportal/helper"
	"simplejobportal/model"
)

type UserRepositoryImpl struct {
	Db *sql.DB
}

func NewUserRepository(Db *sql.DB) UserRepository {
	return &UserRepositoryImpl{Db: Db}
}

// Save implements UserRepository.
func (b *UserRepositoryImpl) Save(ctx context.Context, user model.User) string {
	tx, err := b.Db.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)
	var lastId int
	SQL := "INSERT INTO public.user(username, email, password, user_type) VALUES ($1, $2, $3, $4) RETURNING id"
	err = tx.QueryRowContext(ctx, SQL, user.Username, user.Email, user.Password, user.UserType).Scan(&lastId)
	helper.PanicIfError(err)
	token, err := auth.CreateJWT(lastId, user.UserType)
	helper.PanicIfError(err)
	return token
}
