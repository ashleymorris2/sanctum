package model

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"github.com/google/uuid"
	"strings"
)

// RefreshToken represents a string-based token used for securely refreshing
// authentication sessions.
type RefreshToken string

// NewRefreshToken validates and creates a new RefreshToken instance from the input string.
// The input must adhere to the format `<uuid>.<token>`.
// Returns an error if the input is malformed or invalid.
// This function ensures the token is properly formatted before use in authentication.
func NewRefreshToken(tokenID uuid.UUID, token string) RefreshToken {
	s := fmt.Sprintf("%v.%v", tokenID.String(), token)
	return RefreshToken(s)
}

// ParseRefreshToken validates the format of the input string and converts it into a RefreshToken if valid, or returns an error.
func ParseRefreshToken(s string) (RefreshToken, error) {
	if !isValidRefreshToken(s) {
		return "", fmt.Errorf("invalid refresh token format")
	}
	return RefreshToken(s), nil
}

// Token returns the base64 token part (after the dot)
func (t RefreshToken) Token() string {
	parts := strings.SplitN(string(t), ".", 2)
	return parts[1]
}

// ID returns the UUID part (before the dot)
func (t RefreshToken) ID() uuid.UUID {
	parts := strings.SplitN(string(t), ".", 2)
	return uuid.MustParse(parts[0])
}

// String returns the string representation of the RefreshToken
func (t RefreshToken) String() string {
	return string(t)
}

// Hashed returns the string representation of the RefreshToken but hashed
func (t RefreshToken) Hashed() string {
	hash := sha256.Sum256([]byte(t.Token()))
	return base64.URLEncoding.EncodeToString(hash[:])
}

// isValidRefreshToken ensures token has structure <uuid>.<token>
func isValidRefreshToken(s string) bool {
	parts := strings.SplitN(s, ".", 2)
	if len(parts) != 2 {
		return false
	}
	_, err := uuid.Parse(parts[0])
	return err == nil && len(parts[1]) > 0
}
