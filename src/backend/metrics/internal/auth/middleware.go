package auth

type MiddlewareConfig struct {
	AuthProvider Service
}

func NewMiddlewareConfig(authProvider Service) MiddlewareConfig {
	return MiddlewareConfig{AuthProvider: authProvider}
}
