package handler

import (
	"golang-server/ent"
	"golang-server/internal/models"
	"golang-server/internal/service"
	"golang-server/utils"
	utils_db "golang-server/utils/db"
	"log"

	"github.com/gofiber/fiber/v2"
)

func GetUsers(c *fiber.Ctx) error {
	db, err := utils_db.GetDbFromContext(c)

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Database connection error")
	}

	users := service.GetAllUsers(db)
	return c.JSON(users)
}

func GetUserByID(c *fiber.Ctx) error {
	db, err := utils_db.GetDbFromContext(c)

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Database connection error")
	}

	userId := c.Params("id")
	if userId == "" {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid user ID")
	}

	id := utils.AtoiDefault(userId)

	if id == 0 {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid user ID")
	}
	user, err := service.GetUserByID(db, id)

	if user == nil || err != nil {
		return fiber.NewError(fiber.StatusNotFound, "User not found")
	}

	return c.JSON(user)
}

func CreateUserHandler(c *fiber.Ctx) error {
	var input models.RegisterUserInput

	if err := c.BodyParser(&input); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request")
	}

	db, err := utils_db.GetDbFromContext(c)

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Database connection error")
	}

	if db == nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Database not found in context")
	}

	// Optional: Validate fields (basic example)
	if input.Email == "" || input.Username == "" || input.Password == "" {
		return fiber.NewError(fiber.StatusBadRequest, "Missing required fields")
	}

	user, err := service.CreateUser(c.Context(), db, &input)

	log.Println("User: ", user)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}

func GetMe(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	refreshToken := c.Cookies("refresh_token")

	userId, err := utils.GetUserIdFromTokens(authHeader, refreshToken)
	if err != nil {
		log.Println("Error extracting user ID from tokens:", err)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Could not extract user ID",
		})

	}

	db, err := utils_db.GetDbFromContext(c)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Database connection error",
		})
	}

	var user *ent.User
	user, err = service.GetUserByID(db, utils.AtoiDefault(userId))

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	if user == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	return c.JSON(user)
}
