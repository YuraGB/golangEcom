package service

import (
	"errors"
	"golang-server/ent"
	"golang-server/utils"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	accessSecret  = []byte(utils.GetEnv("JWT_SECRET_ACCESS", "JWT_SECRET_ACCESS"))  
	refreshSecret = []byte(utils.GetEnv("JWT_SECRET_REFRESH", "JWT_SECRET_REFRESH")) 
)

// Custom claims
type TokenClaims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

// Створення access token
func GenerateAccessToken(userID int) (string, error) {
	claims := TokenClaims{
		UserID: strconv.Itoa(int(userID)),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(accessSecret)
}

// Створення refresh token
func GenerateRefreshToken(userID int) (string, error) {
	claims := TokenClaims{
		UserID: strconv.Itoa(int(userID)),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(refreshSecret)
}

// Перевірка access token
func ValidateAccessToken(tokenStr string) (*TokenClaims, error) {
	return parseToken(tokenStr, accessSecret)
}

// Перевірка refresh token
func ValidateRefreshToken(tokenStr string) (*TokenClaims, error) {
	return parseToken(tokenStr, refreshSecret)
}

// Універсальний парсер
func parseToken(tokenStr string, secret []byte) (*TokenClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &TokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		if t.Method.Alg() != jwt.SigningMethodHS256.Alg() {
			return nil, errors.New("unexpected signing method")
		}
		return secret, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*TokenClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}


func GetTokens(user *ent.User) (string, string, error) {
	refreshToken, err := GenerateRefreshToken(user.ID)

	if err != nil {
		return "", "",err
	}

	accessToken, err := GenerateAccessToken(user.ID)

	if err != nil {
		return "", "",err
	}

	return accessToken, refreshToken, nil
}


