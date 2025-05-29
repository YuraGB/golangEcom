package handler

import (
	"context"
	"golang-server/internal/models"
	"golang-server/internal/service"
	"golang-server/utils"
	utils_db "golang-server/utils/db"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
)

func LoginHandler(c *fiber.Ctx) error {
	var userDataInput models.LoginUserInput

	db, err := utils_db.GetDbFromContext(c)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Database connection error"})
	}

	if err := c.BodyParser(&userDataInput); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	loggedInUser, err := service.LoginService(&userDataInput, db)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	// generate refresh token and access tokens
	accessToken, refreshToken, err := service.GetTokens(loggedInUser)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not generate tokens"})
	}

	// set refresh token in cookie
	SetRefreshTokenCookie(c, refreshToken)

	return c.JSON(fiber.Map{
		"user": fiber.Map{
			"id":       loggedInUser.ID,
			"username": loggedInUser.Username,
			"email":    loggedInUser.Email,
			"lastname": loggedInUser.LastName,
			"token":    accessToken,
		},
	})
}

func RegisterUserHandler(c *fiber.Ctx) error {
	db, err := utils_db.GetDbFromContext(c)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Database connection error"})
	}

	var body models.RegisterUserInput

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	hashedPassword, err := utils.HashPassword(body.Password)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not hash password"})
	}

	user := &models.RegisterUserInput{
		Email:    body.Email,
		Password: hashedPassword,
		Username: body.Username,
		LastName: body.LastName,
	}

	createdUser, err := service.CreateUser(context.Background(), db, user)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	// generate refresh token and access tokens
	accessToken, refreshToken, err := service.GetTokens(createdUser)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not generate tokens"})
	}
	// set refresh token in cookie
	SetRefreshTokenCookie(c, refreshToken)

	return c.JSON(fiber.Map{
		"user": fiber.Map{
			"id":       createdUser.ID,
			"username": createdUser.Username,
			"email":    createdUser.Email,
			"lastname": createdUser.LastName,
			"token":    accessToken,
		},
	})
}

func SetRefreshTokenCookie(ctx *fiber.Ctx, token string) {
	isProd := os.Getenv("ENV") == "production"
	cookie := new(fiber.Cookie)
	cookie.Name = "refresh_token"
	cookie.Value = token
	cookie.Expires = time.Now().Add(7 * 24 * time.Hour) // наприклад, 7 днів
	cookie.HTTPOnly = true                              // токен не доступний з JS — захист від XSS
	cookie.Secure = isProd                              // передається тільки через HTTPS
	cookie.SameSite = "Strict"                          // або "Lax", залежить від вимог
	cookie.Path = "/"

	ctx.Cookie(cookie)
}
