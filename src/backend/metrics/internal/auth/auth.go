package auth

import (
	"metrics/internal/db/repositories"
	"metrics/internal/db/sqlc"

	"time"
)

// ByCredentials creates a new instance of CredentialService with the specified database
// queries and JWT secret.
// It configures a default token timeout of 15 minutes, which can be
// modified using functional options.
//
// Example:
//
//	// Create a provider with custom token timeout
//	provider: = auth.ByCredentials(
//	    queries,
//	    []byte("your-jwt-secret"),
//	    auth.WithAuthTokenTTL(24 * time.Hour),
//	)
func ByCredentials(queries *sqlc.Queries, refreshTokenRepo repositories.RefreshTokenRepository, jwtSecret []byte) CredentialService {
	authTokenTTL := 15 * time.Minute       // 15 min
	refreshTokenTTL := 28 * 24 * time.Hour // 28 days

	provider := newCredentialProvider(queries)
	tokenService := newTokenService(jwtSecret, authTokenTTL, refreshTokenTTL, refreshTokenRepo)

	return newCredentialService(provider, tokenService)

}
