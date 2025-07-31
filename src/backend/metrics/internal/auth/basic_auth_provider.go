package auth

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"metrics/internal/db/sqlc"
	"time"
)

type BasicAuthProvider struct {
	queries   *sqlc.Queries
	jwtSecret []byte
}

func NewBasicAuthProvider(queries *sqlc.Queries, jwtSecret []byte) *BasicAuthProvider {
	return &BasicAuthProvider{
		queries:   queries,
		jwtSecret: jwtSecret,
	}
}

func (p *BasicAuthProvider) Register(ctx context.Context, email, password string) (*Result, error) {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	user, err := p.queries.CreateUser(ctx, sqlc.CreateUserParams{
		Email:        email,
		PasswordHash: string(hash),
	})

	if err != nil {
		return nil, err
	}

	token, err := p.generateJWT(user.ID)
	if err != nil {
		return nil, err
	}

	return &Result{UserID: user.ID.String(), Email: user.Email, Token: token}, nil
}

func (p *BasicAuthProvider) Authenticate(ctx context.Context, email, password string) (*Result, error) {
	user, err := p.queries.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return nil, errors.New("invalid email or password")
	}

	token, err := p.generateJWT(user.ID)
	if err != nil {
		return nil, err
	}

	return &Result{UserID: user.ID.String(), Email: user.Email, Token: token}, nil
}

func (p *BasicAuthProvider) generateJWT(userID uuid.UUID) (string, error) {
	claims := jwt.MapClaims{
		"id":  userID.String(),
		"exp": time.Now().Add(15 * time.Minute).Unix(),
		"iat": time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(p.jwtSecret)
}
