package service

import (
	"golang-server/internal/models"
	"golang-server/utils"
	"log"
)

type PostsResponse struct {
	Posts []models.Post `json:"posts"`
}

func GetPosts() ([]models.Post, error) {
	var data PostsResponse

	err := utils.GetDataFromAPI("/posts?limit=6", &data)

	if err != nil {
		log.Printf("Error fetching posts: %v\n", err)
		return nil, err // Handle error if needed
	}

	if data.Posts == nil {
		log.Println("Fetched posts are empty")
	}

	return data.Posts, nil
}