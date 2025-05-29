package utils

import (
	"errors"
	"log"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

var jwtAccessSecret = []byte(getEnv("JWT_SECRET_ACCESS", "your_jwt_secret_key_access"))
var jwtRefreshSecret = []byte(getEnv("JWT_SECRET_REFRESH", "backup_secret_refresh_jwt"))

// getEnv retrieves the value of the environment variable named by the key.
// If the variable is not present, it returns the fallback value.
func getEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return fallback
	}
	return value
}

// GetUserIdFromHeaderToken extracts the user_id from the JWT token in the Authorization header
// and returns it. If the token is invalid or user_id is not found, it returns an error.
func GetUserIdFromHeaderToken(authHeader string) (string, error) {
	if authHeader == "" {
		return "", errors.New("token is empty")
	}

	if !strings.HasPrefix(authHeader, "Bearer ") {
		return "", errors.New("token format is invalid")
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return jwtAccessSecret, nil
	})

	if err != nil || !parsedToken.Valid {
		return "", errors.New("token is invalid")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	log.Println("claims: ", claims)
	if !ok {
		return "", errors.New("failed to parse claims")
	}

	userID, ok := claims["user_id"].(string)
	log.Println("userID: ", userID)
	if !ok {
		return "", errors.New("user_id not found in claims")
	}

	return userID, nil
}

// GetUserIdFromRefreshToken extracts the user_id from the refresh token claims
// and returns it. If the token is invalid or user_id is not found, it returns an error.
func GetUserIdFromRefreshToken(refreshToken string) (string, error) {
	if refreshToken == "" {
		return "", errors.New("refresh token is empty")
	}

	parsedToken, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return jwtRefreshSecret, nil
	})

	if err != nil || !parsedToken.Valid {
		return "", errors.New("refresh token is invalid")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("failed to parse claims")
	}

	userID, ok := claims["user_id"].(string)
	if !ok {
		return "", errors.New("user_id not found in claims")
	}

	return userID, nil
}

func CompareUserIdFromTokens(idFromHeader string, idFromCoockies string) bool {
	if idFromHeader == "" || idFromCoockies == "" {
		return false
	}

	return idFromHeader != idFromCoockies
}

func IsTokenValid(accessToken string, refreshToken string) bool {
	userIdFromHeaderToken, err := GetUserIdFromHeaderToken(accessToken)
	if err != nil {
		return false
	}

	userIdFromRefreshToken, err := GetUserIdFromRefreshToken(refreshToken)
	if err != nil {
		return false
	}
	log.Println("userIdFromHeaderToken: ", userIdFromHeaderToken)
	log.Println("userIdFromRefreshToken: ", userIdFromRefreshToken)
	return CompareUserIdFromTokens(userIdFromHeaderToken, userIdFromRefreshToken)
}

func GetUserIdFromTokens(accessToken string, refreshToken string) (string, error) {

	// refreshToken does not exist in the cookie of the react-native app
	// if IsTokenValid(accessToken, refreshToken) {
	// 	userIdFromHeaderToken, err := GetUserIdFromHeaderToken(accessToken)
	// 	if err != nil {
	// 		return "", err
	// 	}
	// 	return userIdFromHeaderToken, nil
	// }

	userIdFromHeaderToken, err := GetUserIdFromHeaderToken(accessToken)
	log.Println("userIdFromHeaderToken: ", userIdFromHeaderToken)
	log.Println("accessToken ", accessToken)
	if err != nil {
		return "", err
	}

	return userIdFromHeaderToken, nil

}
