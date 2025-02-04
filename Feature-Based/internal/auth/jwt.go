package auth

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type JWTService interface {
	GenerateToken(userID uint) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
	GetUserIDFromToken(token string) (uint, error)
}

type jwtService struct {
	secretKey string
}

func NewJWTService(secretKey string) JWTService {
	return &jwtService{secretKey: secretKey}
}

func (j *jwtService) GenerateToken(userID uint) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(j.secretKey))
}

func (j *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(j.secretKey), nil
	})
}

func (j *jwtService) GetUserIDFromToken(token string) (uint, error) {
	parsedToken, err := j.ValidateToken(token)
	if err != nil {
		return 0, err
	}

	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		if userID, ok := claims["user_id"].(float64); ok {
			return uint(userID), nil
		}
	}
	return 0, errors.New("invalid token")
}
