package service

import "golang-server/ent"

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"lastname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func GetAllUsers(db *ent.Client) []User {
	return []User{
		{ID: 1, Name: "Alice!!!", Email: "test1@emai.com", Password: "1234", LastName: "Smith"},
		{ID: 2, Name: "Bob!!!", Email: "test2@emai.com", Password: "12345"},
	}
}
