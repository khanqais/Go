package user

import (
	"context"
	"errors"
	"fmt"
	"go-auth/internal/auth"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repo      *Repo
	jwtSecret string
}

func NewService(repo *Repo, jwtSecret string) *Service {
	return &Service{repo: repo, jwtSecret: jwtSecret}
}

type RegisterInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type AuthResult struct {
	Token string     `json:"token"`
	User  PublicUser `json:"user"`
}

func (s *Service) Register(ctx context.Context, input RegisterInput) (AuthResult, error) {
	email := strings.ToLower(strings.TrimSpace(input.Email))
	pass := strings.ToLower(strings.TrimSpace(input.Password))
	if email == "" || pass == "" {
		return AuthResult{}, errors.New("email and password are required")
	}
	if len(pass) < 6 {
		return AuthResult{}, errors.New("password is small")
	}
	_, err := s.repo.FindByEmail(ctx, email)
	if err == nil {
		return AuthResult{}, errors.New("email is registered try diff")
	}
	if err != nil && !errors.Is(err, mongo.ErrNilDocument) {
		return AuthResult{}, err
	}
	hashBytes, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return AuthResult{}, fmt.Errorf("Hashing of password falied:%w", err)
	}
	now := time.Now().UTC()

	u := User{
		Email:        email,
		PasswordHash: string(hashBytes),
		Role:         "user",
		CreatedAt:    now,
		UpdatedAt:    now,
	}
	created, err := s.repo.Create(ctx, u)
	if err != nil {
		return AuthResult{}, err
	}
	token, err := auth.CreateToken(s.jwtSecret, created.ID.Hex(), created.Role)
	if err != nil {
		return AuthResult{}, err
	}
	return AuthResult{
		Token: token,
		User:  ToPublic(created),
	}, nil

}
