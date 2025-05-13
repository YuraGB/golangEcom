package models

type Reactions struct {
	Likes    int `json:"likes"`
	Dislikes int `json:"dislikes"`
}

type Post struct {
	ID        int      `json:"id"`
	Title     string   `json:"title"`
	Body      string   `json:"body"`
	Tags      []string `json:"tags"`
	Reactions Reactions `json:"reactions"`
	Views     int      `json:"views"`
	UserID    int      `json:"userId"`
}