package service

import (
	"clother"
	"clother/pkg/repository"
	"database/sql"
	"errors"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user clother.User) (uuid.UUID, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) GetUserByID(id uuid.UUID) (clother.User, error) {
	return s.repo.GetUserByID(id)
}

func (s *AuthService) GetUserByLogin(login string) (clother.User, error) {
	return s.repo.GetUserByLogin(login)
}

func (s *AuthService) GenerateToken(login, password string) (string, error) {
	user, err := s.repo.GetUserByLogin(login)
	if err == sql.ErrNoRows || !checkPasswordHash(user.Password, password) {
		return "", errors.New("wrong login or password")
	} else if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(25 * time.Minute).Unix(),
		Subject:   user.ID.String(),
	})

	return token.SignedString([]byte(os.Getenv("APP_JWT_SECRET")))
}

func (s *AuthService) ParseToken(accessToken string) (uuid.UUID, error) {
	var claims jwt.StandardClaims
	_, err := jwt.ParseWithClaims(accessToken, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("APP_JWT_SECRET")), nil
	})
	if err != nil {
		return uuid.Nil, err
	}

	return uuid.Parse(claims.Subject)
}

func generatePasswordHash(password string) string {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(hashedPassword)
}

func checkPasswordHash(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
