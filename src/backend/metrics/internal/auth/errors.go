package auth

import (
	"errors"
	"net/http"
)

var (
	ErrInvalidJWTToken = NewAPIError(http.StatusUnauthorized, "TOKEN_INVALID", "Access token is invalid")
	ErrExpiredJWTToken = NewAPIError(http.StatusUnauthorized, "TOKEN_EXPIRED", "Access token has expired")
	ErrJWTGeneration   = errors.New("failed to generate JWT token")
	ErrInvalidUserID   = errors.New("userID cannot be nil")

	ErrInvalidRefreshToken = NewAPIError(http.StatusUnauthorized, "REFRESH_INVALID", "Refresh token is invalid")
	ErrRevokedRefreshToken = NewAPIError(http.StatusUnauthorized, "REFRESH_INVALID", "Refresh token is revoked")
	ErrExpiredRefreshToken = errors.New("refresh token has expired")
	ErrRefreshGeneration   = errors.New("failed to generate refresh token")

	ErrInvalidCredentials = NewAPIError(http.StatusUnauthorized, "INVALID_CREDENTIALS", "Invalid credentials")
	ErrHashingFailed      = errors.New("failed to hash password")
	ErrEmptyCredentials   = errors.New("email and password cannot be empty")
	ErrDatabaseError      = errors.New("database error")
)

type APIError struct {
	status  int
	code    string
	message string
}

func (e APIError) Error() string   { return e.message }
func (e APIError) HTTPStatus() int { return e.status }
func (e APIError) Code() string    { return e.code }
func (e APIError) Message() string { return e.message }

func NewAPIError(status int, code, message string) APIError {
	return APIError{status: status, code: code, message: message}
}
