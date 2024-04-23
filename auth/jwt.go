package auth

import (
	"errors"
	"fmt"
	"simplejobportal/config"
	"simplejobportal/helper"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateJWT(userId int, userType int) (string, error) {

	expirationTimeInSecond, secretKey := config.GetJWTData()

	expirationTime := time.Second * time.Duration(expirationTimeInSecond)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":    strconv.Itoa(userId),
		"userType":  strconv.Itoa(userType),
		"expiredAt": time.Now().Add(expirationTime).Unix(),
	})

	tokenString, err := token.SignedString([]byte(secretKey))

	if err != nil {

		return "", err
	}

	return tokenString, nil
}

// ParseJWT parses the JWT token and returns the claims
func ParseJWT(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		// Validate the signing method and return the key
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		_, secretKey := config.GetJWTData()
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, fmt.Errorf("invalid token")
	}
}

// Check if userType is employer
func IsEmployer(userType string) {

	userTypeInt, err := strconv.Atoi(userType)
	helper.PanicIfError(err)
	//Validate if postedByInt != 2 throw error not authorize
	if userTypeInt != 2 {
		// Handle error: Not authorized
		panic(errors.New("not authorized"))
	}
}

// Check if userType is talent
func IsTalent(userType string) {

	userTypeInt, err := strconv.Atoi(userType)
	helper.PanicIfError(err)
	//Validate if postedByInt != 1 throw error not authorize
	if userTypeInt != 1 {
		// Handle error: Not authorized
		panic(errors.New("not authorized"))
	}
}
