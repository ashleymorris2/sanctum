package auth

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"metrics/internal/model"
	"net/http"
	"strings"
	"time"
)

func generateJWT(userID uuid.UUID, tokenTTL time.Duration, jwtSecret []byte) (model.JWTToken, error) {
	if userID == uuid.Nil {
		return "", errors.New("userID cannot be nil")
	}

	now := time.Now()
	claims := jwt.MapClaims{
		"sub": userID.String(),          //Subject - who is the token for
		"exp": now.Add(tokenTTL).Unix(), // Expiration time (unix timestamp)
		"iat": now.Unix(),               // Issued at: time when the token was generated (unix timestamp)
		"nbf": now.Unix(),               //Not before: defines the time before which the JWT cannot be accepted for processing.
		"jti": now.String(),             //JWT ID: an identifier for the JWT, which can be used to prevent the JWT from being replayed.
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return model.NewJWTToken(jwtToken), nil
}

func JwtTokenFromHeader(req *http.Request) (model.JWTToken, error) {
	authHeader := req.Header.Get("Authorization")

	if authHeader != "" && strings.HasPrefix(authHeader, "Bearer ") {
		token := strings.TrimPrefix(authHeader, "Bearer ")
		return model.NewJWTToken(token), nil
	}
	return "", errors.New("no valid Bearer token found")
}
