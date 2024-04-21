package config

import (
	"database/sql"
	"fmt"
	"simplejobportal/helper"

	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

const (
	host                   = "localhost"
	port                   = "5432"
	user                   = "postgres"
	password               = "postgres"
	dbName                 = "simplejobportal"
	jwtExpirationInSeconds = 3600 * 24 * 7
	jwtSecretKey           = "very-secret-key"
)

func DatabaseConnection() *sql.DB {
	sqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)

	db, err := sql.Open("postgres", sqlInfo)
	helper.PanicIfError(err)

	err = db.Ping()
	helper.PanicIfError(err)

	log.Info().Msg("Connected to database!")

	return db
}

func GetJWTData() (int64, string) {

	return jwtExpirationInSeconds, jwtSecretKey
}
