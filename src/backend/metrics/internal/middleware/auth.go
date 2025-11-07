package middleware

import (
	"metrics/internal/auth"

	"metrics/internal/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

// AuthMiddleware creates Echo middleware for JWT authentication
// Accepts any auth service that can validate tokens (CredentialService, OAuthService, APIKeyService)
func AuthMiddleware(config auth.MiddlewareConfig) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			jwtToken, err := auth.JWTFromHeader(c.Request())
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "Missing or invalid authorization header")
			}

			if userID, ok := validateAccessToken(jwtToken, config); ok {
				// Set userID in context for use in handlers
				c.Set("userID", userID)
				return next(c)
			}

			return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
		}
	}
}

func validateAccessToken(token model.JWTToken, config auth.MiddlewareConfig) (userID string, valid bool) {
	claims, err := config.AuthProvider.ValidateToken(model.JWTToken(token))
	if err != nil {
		return "", false
	}

	subject, err := claims.GetSubject()
	if err != nil {
		return "", false
	}

	return subject, true
}
