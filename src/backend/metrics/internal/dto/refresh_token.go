package dto

// RefreshTokenRequest represents the payload for a token refresh
type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

// RefreshTokenResponse represents the response after a successful token refresh
type RefreshTokenResponse struct {
	AuthToken       string  `json:"auth_token"`
	RefreshToken    string  `json:"refresh_token"`
	RefreshTokenTTL float64 `json:"refreshTokenTTL"`
	UserId          string  `json:"userId"`
}
