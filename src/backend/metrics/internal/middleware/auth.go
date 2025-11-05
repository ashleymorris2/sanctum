package middleware

import (
	"metrics/internal/auth"
	"metrics/internal/auth/tokens"
	"metrics/internal/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func AuthMiddleware(config auth.MiddlewareConfig) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			token, _ := tokens.JwtTokenFromHeader(c.Request())
			if userID, ok := validateAccessToken(token.String(), config); ok {
				// Set userID in context for use in handler
				c.Set("userID", userID)
				return next(c)
			}

			return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
		}
	}
}

func validateAccessToken(token string, config auth.MiddlewareConfig) (userID string, valid bool) {
	credentialAuth, ok := config.AuthProvider.(*auth.CredentialService)
	if !ok {
		return "", false
	}

	claims, err := credentialAuth.ValidateJwtToken(model.JWTToken(token))
	if err != nil {
		return "", false
	}

	subject, err := claims.GetSubject()
	if err != nil {
		return "", false
	}

	return subject, true
}
