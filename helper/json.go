package helper

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

func ReadRequestBody(r *http.Request, result interface{}) {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(result)
	PanicIfError(err)
}

func WriteResponseBody(write http.ResponseWriter, response interface{}) {
	write.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(write)

	err := encoder.Encode(response)
	PanicIfError(err)
}

func ReadBearerToken(r *http.Request) (string, error) {
	// Get Bearer Token from Header
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		// Handle missing Authorization header
		return "", errors.New("authorization header is missing")
	}

	// Check if the Authorization header starts with "Bearer "
	if !strings.HasPrefix(authHeader, "Bearer ") {
		return "", errors.New("invalid authorization header format")
	}

	// Extract the token from the Authorization header
	tokenString := authHeader[len("Bearer "):]

	return tokenString, nil
}
