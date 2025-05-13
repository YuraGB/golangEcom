package service

import (
	"context"
	"errors"
	"golang-server/ent"
	"golang-server/ent/user"
	"golang-server/internal/models"
	"log"
)

func GetAllUsers(db *ent.Client) []ent.User {
	return []ent.User{}
}

func CreateUser(ctx context.Context, client *ent.Client, input *models.RegisterUserInput) (*ent.User, error) {
	if input == nil {
		log.Println("RegisterUserInput is nil")
		return nil, errors.New("invalid input: nil")
	}
	log.Println("Creating user with input:", input.Email)

	user, err := client.User.
		Create().
		SetUsername(input.Username).
		SetEmail(input.Email).
		SetPassword(input.Password).
		SetLastName(input.LastName).
		Save(context.Background())

	if err != nil {
		log.Println("Failed to create user:", err)
		return nil, err
	}

	return user, nil
}

func GetUserByID(db *ent.Client, id int) (*ent.User, error) {
	user, err := db.User.Query().Where(user.ID(id)).Only(context.Background())
	if err != nil {
		return nil, err
	}
	return user, nil
}
