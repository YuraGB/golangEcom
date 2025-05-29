package utils

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func GetUserIdFromCtx(c *fiber.Ctx) (int, error) {
	authHeader := c.Get("Authorization")
	userIdStr, err := GetUserIdFromHeaderToken(authHeader)
	log.Println("UserIdStr", userIdStr)
	if err != nil {
		return 0, fiber.NewError(fiber.StatusUnauthorized, "Invalid access token")
	}
	return AtoiDefault(userIdStr), nil
}
