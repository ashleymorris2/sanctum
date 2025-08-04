package models

type RefreshToken string

func NewRefreshToken(s string) RefreshToken {
	return RefreshToken(s)
}

func (t RefreshToken) String() string {
	return string(t)
}
