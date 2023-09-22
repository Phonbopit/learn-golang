package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

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
			render.JSON(w, r, post)
			return
		}
	}

	render.JSON(w, r, Error{Message: fmt.Sprintf("PostID %s not found", postID)})
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

	render.JSON(w, r, Error{Message: fmt.Sprintf("UserID %s not found", userID)})
}
