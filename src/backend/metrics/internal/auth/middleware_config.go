package auth

type MiddlewareConfig struct {
	AuthProvider Provider
}

func NewMiddlewareConfig(authProvider Provider) MiddlewareConfig {
	return MiddlewareConfig{AuthProvider: authProvider}
}
