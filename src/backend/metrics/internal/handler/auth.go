package handler

import (
	"errors"
	"github.com/labstack/echo/v4"
	"metrics/internal/auth"
	"metrics/internal/dto"
	"metrics/internal/model"
	"metrics/internal/validators"
	"net/http"
)

const secureCookie = false

type Auth struct {
	authProvider auth.Provider
}

func NewAuth(authProvider auth.Provider) *Auth {
	return &Auth{authProvider: authProvider}
}

func (a *Auth) Login(c echo.Context) error {
	ctx := c.Request().Context()

	var req dto.LoginRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, ErrorResponse{Message: "Invalid request"})
	}

	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, validators.FormatErrors(err))
	}

	authResult, err := a.authProvider.Authenticate(ctx, req.Email, req.Password)
	if err != nil {
		if errors.Is(err, auth.ErrAuthFailure) {
			return echo.NewHTTPError(http.StatusUnauthorized, ErrorResponse{
				Message: "Invalid credentials",
			})
		} else {
			return echo.NewHTTPError(http.StatusInternalServerError, ErrorResponse{
				Message: "Internal server error",
			})
		}
	}

	return c.JSON(http.StatusOK, dto.LoginResponse{
		AuthToken:       authResult.AuthToken.String(),
		RefreshToken:    authResult.RefreshToken.String(),
		RefreshTokenTTL: authResult.RefreshTokenTTL.Seconds(),
		UserId:          authResult.UserID,
	})
}

func (a *Auth) VerifyAuthToken(c echo.Context) error {
	token, err := auth.JwtTokenFromHeader(c.Request())
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Missing or invalid token")
	}

	credentialService, ok := a.authProvider.(*auth.CredentialService)
	if !ok {
		//Not configured for credential auth
		return echo.NewHTTPError(http.StatusInternalServerError, "`Internal server error`")
	}

	claims, err := credentialService.ValidateJwtToken(token)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid or expired token")
	}

	sub, err := claims.GetSubject()
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
	}

	return c.JSON(http.StatusOK, map[string]string{
		"id": sub,
	})
}

func (a *Auth) RefreshAuthToken(c echo.Context) error {

	cookie, err := c.Cookie("refresh_token")
	if err != nil {
		return c.JSON(http.StatusUnauthorized, "Invalid refresh token")
	}

	credentialService, ok := a.authProvider.(*auth.CredentialService)
	if !ok {
		//Not configured for credential auth
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}

	tokenPair, err := credentialService.RefreshJwtToken(c.Request().Context(), model.RefreshToken(cookie.Value))
	if err != nil {
		// Error handling...
		return echo.NewHTTPError(http.StatusUnauthorized, "invalid token")
	}

	return c.JSON(http.StatusOK, dto.RefreshTokenResponse{
		AuthToken: tokenPair.AuthToken.String(),
		UserId:    tokenPair.UserID,
	})
}
