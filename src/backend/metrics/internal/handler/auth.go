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

	c.SetCookie(&http.Cookie{
		Name:     "refresh_token",
		Value:    authResult.RefreshToken.Token(),
		Path:     "/", // only sent on /refresh route
		HttpOnly: true,
		Secure:   secureCookie, // only send over HTTPS
		SameSite: http.SameSiteNoneMode,
		MaxAge:   int(authResult.RefreshTokenTTL.Seconds()),
	})

	return c.JSON(http.StatusOK, dto.LoginResponse{
		AuthToken:       authResult.AuthToken.String(),
		RefreshToken:    authResult.RefreshToken.Token(),
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

	claims, err := credentialService.ValidateJwtToken(model.JWTToken(token))
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
	var req dto.RefreshTokenRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request")
	}

	if err := c.Validate(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	credentialService, ok := a.authProvider.(*auth.CredentialService)
	if !ok {
		//Not configured for credential auth
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}

	refreshToken, err := model.ParseRefreshToken(req.RefreshToken)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid refresh token")
	}

	tokenPair, err := credentialService.RefreshJwtToken(c.Request().Context(), refreshToken)
	if err != nil {
		// Error handling...
		return echo.NewHTTPError(http.StatusUnauthorized, "invalid token")
	}

	c.SetCookie(&http.Cookie{
		Name:     "refresh_token",
		Value:    tokenPair.RefreshToken.Token(),
		Path:     "/", // only sent on /refresh route
		HttpOnly: true,
		Secure:   secureCookie, // only send over HTTPS
		SameSite: http.SameSiteNoneMode,
		MaxAge:   int(tokenPair.RefreshTokenTTL.Seconds()),
	})

	return c.JSON(http.StatusOK, dto.RefreshTokenResponse{
		AuthToken: tokenPair.AuthToken.String(),
	})
}
