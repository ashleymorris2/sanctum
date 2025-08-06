package dto

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=128"`
}

type LoginResponse struct {
	AuthToken       string  `json:"authToken"`
	UserId          string  `json:"userId"`
	RefreshToken    string  `json:"refreshToken"`
	RefreshTokenTTL float64 `json:"refreshTokenTTL"`
}
