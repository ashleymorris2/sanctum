package models

type JWTToken string

func NewJWTToken(s string) JWTToken {
	return JWTToken(s)
}

func (t JWTToken) String() string {
	return string(t)
}
