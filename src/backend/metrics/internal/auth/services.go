package auth

import (
	"context"
	"metrics/internal/model"

	"github.com/golang-jwt/jwt/v5"
)

// CredentialService handles user authentication and registration using basic email-password credentials.
// It hashes passwords using bcrypt, interacts with the database to persist user records, and generates JWT tokens
// for session management. It relies on sqlc.Queries for database operations and supports configurable JWT timeouts.
type credentialService struct {
	provider     credentialAuthProvider
	tokenService *tokenService
}

func newCredentialService(provider credentialAuthProvider, tokenService *tokenService) CredentialService {
	return &credentialService{
		provider:     provider,
		tokenService: tokenService,
	}
}

func (s *credentialService) Register(ctx context.Context, credentials EmailPasswordCredentials) (*SessionResult, error) {
	authResult, err := s.provider.registerWithCredentials(ctx, credentials)
	if err != nil {
		return nil, err
	}

	tokenPair, err := s.tokenService.generateTokenPair(ctx, authResult.userID)
	if err != nil {
		return nil, err
	}

	return &SessionResult{
		UserID:    authResult.userID.String(),
		Email:     authResult.email,
		TokenPair: tokenPair,
	}, nil
}

func (s *credentialService) Login(ctx context.Context, credentials EmailPasswordCredentials) (*SessionResult, error) {
	authResult, err := s.provider.authenticateWithCredentials(ctx, credentials)
	if err != nil {
		return nil, err
	}

	tokenPair, err := s.tokenService.generateTokenPair(ctx, authResult.userID)
	if err != nil {
		return nil, err
	}

	return &SessionResult{
		UserID:    authResult.userID.String(),
		Email:     authResult.email,
		TokenPair: tokenPair,
	}, nil
}

func (s *credentialService) RefreshSession(ctx context.Context, refreshToken model.RefreshToken) (*TokenPair, error) {
	tokenPair, err := s.tokenService.renewTokenPair(ctx, refreshToken)
	if err != nil {
		return nil, err
	}

	_, err = s.tokenService.validateJWT(tokenPair.AccessToken)
	if err != nil {
		return nil, err
	}

	return tokenPair, nil
}

func (s *credentialService) ValidateToken(token model.JWTToken) (jwt.MapClaims, error) {
	claims, err := s.tokenService.validateJWT(token)
	if err != nil {
		return nil, err
	}

	//subject, err := claims.GetSubject()
	//if err != nil {
	//	return uuid.Nil, ErrInvalidJWTToken
	//}
	//
	//userID, err := uuid.Parse(subject)
	//if err != nil {
	//	return uuid.Nil, ErrInvalidJWTToken
	//}

	return claims, nil
}

func (s *credentialService) Logout(ctx context.Context, refreshToken model.RefreshToken) error {
	return s.tokenService.revokeRefreshToken(ctx, refreshToken)
}
