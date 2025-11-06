package auth

import (
	"context"
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"metrics/internal/db/repositories"
	"metrics/internal/model"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type tokenService struct {
	jwtSecret        []byte
	jwtTTL           time.Duration
	refreshTTL       time.Duration
	refreshTokenRepo repositories.RefreshTokenRepository
}

// NewTokenService creates a new token service instance
func newTokenService(
	jwtSecret []byte,
	jwtTTL time.Duration,
	refreshTTL time.Duration,
	refreshTokenRepo repositories.RefreshTokenRepository,
) *tokenService {
	return &tokenService{
		jwtSecret:        jwtSecret,
		jwtTTL:           jwtTTL,
		refreshTTL:       refreshTTL,
		refreshTokenRepo: refreshTokenRepo,
	}
}

// generateJWT creates a new JWT token for the given user ID
func (m *tokenService) generateJWT(userID uuid.UUID) (model.JWTToken, error) {
	if userID == uuid.Nil {
		return "", ErrInvalidUserID
	}

	now := time.Now()
	claims := jwt.MapClaims{
		"sub": userID.String(),          //Subject - who is the token for
		"exp": now.Add(m.jwtTTL).Unix(), // Expiration time (unix timestamp)
		"iat": now.Unix(),               // Issued at: time when the token was generated (unix timestamp)
		"nbf": now.Unix(),               //Not before: defines the time before which the JWT cannot be accepted for processing.
		"jti": now.String(),             //JWT ID: an identifier for the JWT, which can be used to prevent the JWT from being replayed.
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtToken, err := token.SignedString(m.jwtSecret)
	if err != nil {
		return "", ErrJWTGeneration
	}

	return model.NewJWTToken(jwtToken), nil
}

// validateJWT parses and validates a JWT token, returning its claims
func (m *tokenService) validateJWT(jwtToken model.JWTToken) (jwt.MapClaims, error) {
	// Parse the token
	token, err := jwt.Parse(jwtToken.String(), func(token *jwt.Token) (interface{}, error) {
		// Validate the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidJWTToken
		}
		return m.jwtSecret, nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, ErrExpiredJWTToken
		}
		return nil, ErrInvalidJWTToken
	}

	// Extract and validate claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, ErrInvalidJWTToken
}

// generateRefreshToken creates a new cryptographically secure refresh token
func (m *tokenService) generateRefreshToken() (model.RefreshToken, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	return model.NewRefreshToken(uuid.New(), base64.URLEncoding.EncodeToString(b)), nil
}

// storeRefreshToken persists a refresh token for a user
func (m *tokenService) storeRefreshToken(ctx context.Context, token model.RefreshToken, userID uuid.UUID) error {
	return m.refreshTokenRepo.InsertRefreshToken(ctx, token, userID, m.refreshTTL)
}

func (m *tokenService) validateRefreshToken(ctx context.Context, refreshToken model.RefreshToken) (*refreshTokenInfo, error) {
	// Retrieve the refresh token
	storedToken, err := m.refreshTokenRepo.GetRefreshToken(ctx, refreshToken)
	if err != nil {
		return nil, ErrInvalidRefreshToken
	}

	if subtle.ConstantTimeCompare([]byte(storedToken.Token), []byte(refreshToken.Hashed())) != 1 {
		return nil, ErrInvalidRefreshToken
	}

	// Check if the token is revoked
	if storedToken.Revoked {
		return nil, ErrRevokedRefreshToken
	}

	// Check if the token has expired
	if time.Now().After(storedToken.ExpiresAt) {
		return nil, ErrExpiredRefreshToken
	}

	return &refreshTokenInfo{
		userID:    storedToken.UserID,
		expiresAt: storedToken.ExpiresAt,
	}, nil
}

// generateTokenPair creates both a JWT and refresh token for a user
func (m *tokenService) generateTokenPair(ctx context.Context, userID uuid.UUID) (*TokenPair, error) {
	// Generate JWT
	jwtToken, err := m.generateJWT(userID)
	if err != nil {
		return nil, ErrRefreshGeneration
	}

	// Generate refresh token
	refreshToken, err := m.generateRefreshToken()
	if err != nil {
		return nil, err
	}

	// Store refresh token
	if err := m.storeRefreshToken(ctx, refreshToken, userID); err != nil {
		return nil, err
	}

	return &TokenPair{
		AccessToken:     jwtToken,
		AccessTokenTTL:  m.jwtTTL,
		RefreshToken:    refreshToken,
		RefreshTokenTTL: m.refreshTTL,
	}, nil
}

// renewTokenPair validates a refresh token and generates a new JWT if valid
func (m *tokenService) renewTokenPair(ctx context.Context, refreshToken model.RefreshToken) (*TokenPair, error) {
	// Validate refresh token
	info, err := m.validateRefreshToken(ctx, refreshToken)
	if err != nil {
		return nil, err
	}

	// Generate new JWT
	jwtToken, err := m.generateJWT(info.userID)
	if err != nil {
		return nil, err
	}

	return &TokenPair{
		AccessToken:     jwtToken,
		AccessTokenTTL:  m.jwtTTL,
		RefreshToken:    refreshToken, // Reuse the same refresh token
		RefreshTokenTTL: m.refreshTTL,
	}, nil
}

// revokeRefreshToken marks a refresh token as invalid
func (m *tokenService) revokeRefreshToken(ctx context.Context, refreshToken model.RefreshToken) error {
	return m.refreshTokenRepo.RevokeRefreshToken(ctx, refreshToken)
}

type TokenValidator struct {
	tokenService *tokenService
}

// NewTokenValidator creates a public token validator
// This is exported so it can be used independently of auth services
func NewTokenValidator(jwtSecret []byte) *TokenValidator {
	// Create a minimal token service just for validation
	ts := &tokenService{
		jwtSecret: jwtSecret,
	}
	return &TokenValidator{tokenService: ts}
}

// ValidateToken validates a JWT token and returns the user ID
func (v *TokenValidator) ValidateToken(token model.JWTToken) (uuid.UUID, error) {
	claims, err := v.tokenService.validateJWT(token)
	if err != nil {
		return uuid.Nil, err
	}

	return uuid.Parse(claims["sub"].(string))
}

// ValidateTokenFromHeader extracts and validates a JWT from an HTTP request
func (v *TokenValidator) ValidateTokenFromHeader(req *http.Request) (uuid.UUID, error) {
	token, err := jwtFromHeader(req)
	if err != nil {
		return uuid.Nil, err
	}
	return v.ValidateToken(token)
}

// jwtFromHeader extracts a JWT token from the Authorization header
func jwtFromHeader(req *http.Request) (model.JWTToken, error) {
	authHeader := req.Header.Get("Authorization")

	if authHeader != "" && strings.HasPrefix(authHeader, "Bearer ") {
		token := strings.TrimPrefix(authHeader, "Bearer ")
		return model.NewJWTToken(token), nil
	}

	return "", errors.New("no valid Bearer token found in Authorization header")
}
