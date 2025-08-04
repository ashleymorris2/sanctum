package middleware

import (
	"github.com/labstack/echo/v4"
	"metrics/internal/auth"
	"metrics/internal/models"
	"net/http"
	"strings"
)

func AuthMiddleware(config auth.MiddlewareConfig) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// 1. Extract access token from the Authorization header
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader != "" && strings.HasPrefix(authHeader, "Bearer ") {
				token := strings.TrimPrefix(authHeader, "Bearer ")

				if userID, ok := validateAccessToken(token, config); ok {
					// Set userID in context for use in handlers
					c.Set("userID", userID)
					return next(c)
				}
			}

			// 2. If access token is not valid, check refresh token (cookie example)
			refreshCookie, err := c.Cookie("refresh_token")
			if err == nil {
				if userID, jwtToken, ok := refreshAccessToken(refreshCookie.Value, config); ok {
					c.Set("userID", userID)
					issueNewTokensHeaders(c, jwtToken.String(), refreshCookie.Value)
					return next(c)
				}
			}

			// 3. Unauthorized if both tokens invalid
			return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
		}
	}
}

func validateAccessToken(token string, config auth.MiddlewareConfig) (userID string, valid bool) {
	credentialAuth, ok := config.AuthProvider.(*auth.CredentialService)
	if !ok {
		return "", false
	}
	claims, err := credentialAuth.ValidateJwtToken(models.JWTToken(token))
	if err != nil {
		return "", false
	}
	subject, err := claims.GetSubject()
	if err != nil {
		return "", false
	}

	return subject, true
}

func refreshAccessToken(refreshToken string, config auth.MiddlewareConfig) (userID string, jwtToken models.JWTToken, valid bool) {
	credentialAuth, ok := config.AuthProvider.(*auth.CredentialService)
	if !ok {
		return "", "", false
	}

	jwtToken, err := credentialAuth.RefreshJwtToken(models.RefreshToken(refreshToken))
	if err != nil {
		return "", "", false
	}

	claims, err := credentialAuth.ValidateJwtToken(jwtToken)
	if err != nil {
		return "", "", false
	}

	subject, err := claims.GetSubject()
	if err != nil {
		return "", "", false
	}

	return subject, jwtToken, true
}

func issueNewTokensHeaders(c echo.Context, accessToken string, refreshToken string) {
	c.Response().Header().Set("X-Access-Token", accessToken)
	c.Response().Header().Set("X-Refresh-Token", refreshToken)
}
