package main

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

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

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	r.Get("/posts", FindAllPosts)
	r.Post("/posts", CreateNewPost)
	r.Get("/posts/{postID}", FindPostByID)
	r.Get("/users", FindAllUsers)
	r.Get("/users/{userID}", FindUserByID)

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 - Not Found"))
	})

	r.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("405 - Method Not Allowed"))
	})

	http.ListenAndServe(":9998", r)
}

func FindAllPosts(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, posts)
}

func CreateNewPost(w http.ResponseWriter, r *http.Request) {
	var post Post
	err := render.DecodeJSON(r.Body, &post)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 - Bad Request"))
		return
	}

	post.ID = strconv.Itoa(len(posts) + 1)

	posts = append(posts, post)
	render.JSON(w, r, post)
}

func FindAllUsers(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, users)
}

func FindPostByID(w http.ResponseWriter, r *http.Request) {
	postID := chi.URLParam(r, "postID")

	var post Post

	for _, p := range posts {
		if p.ID == postID {
			post = p
		}
	}

	render.JSON(w, r, post)
}

func FindUserByID(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")

	var user User

	for _, u := range users {
		if u.ID == userID {
			user = u
			render.JSON(w, r, user)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("User Not Found"))
}
