package main

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Active   bool   `json:"active"`
}

type Post struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Body     string `json:"body"`
	AuthorID int    `json:"author_id"`
}

type Error struct {
	Message string `json:"message"`
}

var users = []User{
	{ID: "1", Name: "John Doe", Username: "johndoe", Active: true},
	{ID: "2", Name: "Jane Doe", Username: "janedoe", Active: true},
	{ID: "3", Name: "Michael Jordan", Username: "michaeljordan", Active: true},
	{ID: "4", Name: "John Smith", Username: "johnsmith", Active: true},
}

var posts = []Post{
	{ID: "1", Title: "Hello World", Body: "This is my first post", AuthorID: 1},
	{ID: "2", Title: "Hello World 2", Body: "This is my second post", AuthorID: 1},
	{ID: "3", Title: "Hello World 3", Body: "This is my third post", AuthorID: 2},
	{ID: "4", Title: "Hello World 4", Body: "This is my fourth post", AuthorID: 3},
}
