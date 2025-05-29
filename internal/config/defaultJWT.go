package config

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var accessSecret = []byte("your_jwt_secret_key_access") // заміни на свій ключ

func DebugJWTMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")

		if authHeader == "" {
			fmt.Println("Authorization header not found")
			return c.Next()
		}

		tokenStr := authHeader[len("Bearer "):]

		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			return accessSecret, nil
		})

		if err != nil {
			fmt.Println("JWT Parse Error:", err)
			return c.Next()
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			expUnix := int64(claims["exp"].(float64))
			iatUnix := int64(claims["iat"].(float64))

			expTime := time.Unix(expUnix, 0).UTC()
			iatTime := time.Unix(iatUnix, 0).UTC()
			current := time.Now().UTC()

			fmt.Println("✅ JWT Debug Info: ", tokenStr)
			fmt.Println("- iat     :", iatTime)
			fmt.Println("- exp     :", expTime)
			fmt.Println("- now     :", current)
			fmt.Println("- expired :", current.After(expTime))
		} else {
			fmt.Println("Invalid JWT claims or token")
		}

		return c.Next()
	}
}
