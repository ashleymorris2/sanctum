package auth

import (
	"metrics/internal/db/sqlc"
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
//	    []byte("jwt-secret"),
//	)
func ByCredentials(queries *sqlc.Queries, tokenService *TokenService) CredentialService {
	provider := newCredentialProvider(queries)
	return newCredentialService(provider, tokenService)
}
