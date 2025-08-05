package dto

import "metrics/internal/model"

type CredentialAuthResult struct {
	UserID       string
	JWTToken     model.JWTToken
	RefreshToken model.RefreshToken
	Email        string
}
