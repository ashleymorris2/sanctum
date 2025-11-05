package auth

import "errors"

var (
	ErrInvalidJWTToken = errors.New("invalid JWT token")
	ErrExpiredJWTToken = errors.New("JWT token has expired")
	ErrJWTGeneration   = errors.New("failed to generate JWT token")
	ErrInvalidUserID   = errors.New("userID cannot be nil")

	ErrInvalidRefreshToken = errors.New("invalid refresh token")
	ErrExpiredRefreshToken = errors.New("refresh token has expired")
	ErrRevokedRefreshToken = errors.New("refresh token has been revoked")
	ErrRefreshGeneration   = errors.New("failed to generate refresh token")

	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrHashingFailed      = errors.New("failed to hash password")
	ErrEmptyCredentials   = errors.New("email and password cannot be empty")
	ErrDatabaseError      = errors.New("database error")
)
