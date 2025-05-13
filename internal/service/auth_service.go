package service

import (
	"context"
	"errors"
	"golang-server/ent"
	"golang-server/ent/user"
	"golang-server/internal/models"
	"golang-server/utils"
	"log"
)

func LoginService(userInput *models.LoginUserInput, db *ent.Client) (*ent.User, error) {
	u, err := db.User.
		Query().
		Where(user.EmailEQ(userInput.Email)).
		Only(context.Background())

	if err != nil {
		log.Println("Error in LoginService: Can't get user", err.Error())
		return nil, errors.New("користувача не знайдено")
	}

	isValid := utils.CompareHashAndPassword(u.Password, userInput.Password)

	if !isValid {
		log.Println("Error in LoginService: Password not verified. userId:", u.ID)
		return nil, errors.New("неправильний пароль")
	}

	return u, nil
}

func EditService(userData *ent.User) ent.User {
	return ent.User{
		ID:       userData.ID,
		Username:     userData.Username,
		Email:    userData.Email,
		LastName: userData.LastName,
	}
}

func DeleteService(id int) bool {
	// Simulate a user deletion process
	// In a real application, you would delete the user from a database
	return true
}
