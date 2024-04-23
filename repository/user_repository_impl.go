package repository

import (
	"context"
	"database/sql"
	"errors"
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

// GetByEmail implements UserRepository.
func (b *UserRepositoryImpl) GetByEmail(ctx context.Context, email string) (model.User, error) {
	tx, err := b.Db.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	SQL := `SELECT
			a.id,
			a.username,
			a.email,
			a.password,
			a.user_type
			FROM public.user a
			WHERE a.email = $1`

	result, errQuery := tx.QueryContext(ctx, SQL, email)
	helper.PanicIfError(errQuery)

	defer result.Close()

	user := model.User{}

	if result.Next() {
		err := result.Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.UserType)
		helper.PanicIfError(err)
		return user, err
	} else {
		return user, errors.New("email not found")
	}
}

// Save implements UserRepository.
func (b *UserRepositoryImpl) Save(ctx context.Context, user model.User) (string, error) {
	tx, err := b.Db.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	SQL := `SELECT
			a.id,
			a.username,
			a.email,
			a.password,
			a.user_type
			FROM public.user a
			WHERE a.email = $1 AND a.user_type = $2`

	result, errQuery := tx.QueryContext(ctx, SQL, user.Email, user.UserType)
	helper.PanicIfError(errQuery)

	defer result.Close()

	if result.Next() {
		return "", errors.New("email has already used for this role")
	}

	var lastId int
	SQL = "INSERT INTO public.user(username, email, password, user_type) VALUES ($1, $2, $3, $4) RETURNING id"
	err = tx.QueryRowContext(ctx, SQL, user.Username, user.Email, user.Password, user.UserType).Scan(&lastId)
	helper.PanicIfError(err)
	token, err := auth.CreateJWT(lastId, user.UserType)
	helper.PanicIfError(err)
	return token, nil
}
