package auth

type Provider interface {
	Authenticate(email, password string) (Result, error)
}

type Request struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=128"`
}

type Result struct {
	UserID string
	Token  string
	Email  string
}
