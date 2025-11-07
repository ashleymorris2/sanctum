package handler

import (
	"errors"
	"metrics/internal/auth"

	"metrics/internal/dto"
	"metrics/internal/model"
	"metrics/internal/validators"
	"net/http"

	"github.com/labstack/echo/v4"
)

const secureCookie = false

type AuthHandler struct {
	authProvider auth.CredentialService
}

func NewAuthHandler(authProvider auth.CredentialService) *AuthHandler {
	return &AuthHandler{authProvider: authProvider}
}

// Login authenticates a user and returns access and refresh tokens
//
//	@Summary		User login
//	@Description	Authenticate a user with their email and password
//	@Tags			authentication
//	@Accept			json
//	@Produce		json
//	@Param			request	body		dto.LoginRequestDoc	true	"Login credentials"
//	@Success		200		{object}	dto.LoginResponse
//	@Failure		400		{object}	ErrorResponse
//	@Failure		401		{object}	ErrorResponse
//	@Failure		500		{object}	ErrorResponse
//	@Router			/auth/login [post]
func (a *AuthHandler) Login(c echo.Context) error {
	ctx := c.Request().Context()

	var req dto.LoginRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, ErrorResponse{Message: "Invalid request"})
	}

	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, validators.FormatErrors(err))
	}

	authResult, err := a.authProvider.Login(ctx, auth.EmailPasswordCredentials{
		Email:    req.Email,
		Password: req.Password,
	})

	if err != nil {
		if errors.Is(err, auth.ErrInvalidCredentials) {
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
		AuthToken:       authResult.TokenPair.AccessToken.String(),
		RefreshToken:    authResult.TokenPair.RefreshToken.String(),
		RefreshTokenTTL: authResult.TokenPair.RefreshTokenTTL.Seconds(),
		UserId:          authResult.UserID,
	})
}

// VerifyAuthToken
//
//	@Summary					User login
//	@Description				Authenticate a user with their email and password
//	@securityDefinitions.basic	BasicAuth
//	@Tags						authentication
//	@Accept						json
//	@Produce					json
//	@Success					200	{object}	dto.LoginResponse
//	@Failure					400	{object}	ErrorResponse
//	@Failure					401	{object}	ErrorResponse
//	@Failure					500	{object}	ErrorResponse
//	@Router						/auth/verify [post]
func (a *AuthHandler) VerifyAuthToken(c echo.Context) error {
	token, err := auth.JWTFromHeader(c.Request())
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Missing or invalid token")
	}

	claims, err := a.authProvider.ValidateToken(token)
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

// RefreshAuthToken
//
//	@Summary					User login
//	@Description				Authenticate a user with their email and password
//	@securityDefinitions.basic	BasicAuth
//	@Tags						authentication
//	@Accept						json
//	@Produce					json
//	@Success					200	{object}	dto.LoginResponse
//	@Failure					400	{object}	ErrorResponse
//	@Failure					401	{object}	ErrorResponse
//	@Failure					500	{object}	ErrorResponse
//	@Router						/auth/refresh [post]
func (a *AuthHandler) RefreshAuthToken(c echo.Context) error {
	cookie, err := c.Cookie("refresh_token")
	if err != nil {
		return c.JSON(http.StatusUnauthorized, "Invalid refresh token")
	}

	tokenPair, err := a.authProvider.RefreshSession(c.Request().Context(), model.RefreshToken(cookie.Value))
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "invalid token")
	}

	return c.JSON(http.StatusOK, dto.RefreshTokenResponse{
		AccessToken: tokenPair.AccessToken.String(),
	})
}
