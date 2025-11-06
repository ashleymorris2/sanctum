package auth

type MiddlewareConfig struct {
	AuthProvider CredentialService
}

func NewMiddlewareConfig(authProvider CredentialService) MiddlewareConfig {
	return MiddlewareConfig{AuthProvider: authProvider}
}
