package service

import (
	"time"

	"Filo.Hack/internal/app/middleware"
	"Filo.Hack/internal/app/repository"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	residentRepo *repository.ResidentRepository
	secretKey    string
}

func NewAuthService(residentRepo *repository.ResidentRepository, secretKey string) *AuthService {
	return &AuthService{
		residentRepo: residentRepo,
		secretKey:    secretKey,
	}
}

func (s *AuthService) SignIn(email, password string) (string, error) {
	resident, err := s.residentRepo.GetResidentByEmail(email)

	// Проверка пароля
	err = bcrypt.CompareHashAndPassword([]byte(resident.Password), []byte(password))
	if err != nil {
		return "", err
	}

	// Генерация JWT токена с email и ролью
	claims := &middleware.CustomClaims{
		Email: resident.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(s.secretKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
