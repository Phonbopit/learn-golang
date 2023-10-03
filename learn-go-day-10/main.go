package main

import (
	"fmt"
	"net/http"

	"database/sql"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	initDB()

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
	rows, err := DB.Query("SELECT * FROM posts")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	var posts []Post

	for rows.Next() {
		var post Post

		err := rows.Scan(&post.ID, &post.Title, &post.Body)
		if err != nil {
			panic(err.Error())
		}

		posts = append(posts, post)
	}

	if err := rows.Err(); err != nil {
		panic(err.Error())
	}

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

	query := "INSERT INTO `posts` (`title`, `body`, `created_at`) VALUES (?, ?, NOW())"
	insertResult, err := DB.Exec(query, post.Title, post.Body)

	if err != nil {
		panic(err.Error())
	}

	fmt.Println("insertResult: ", insertResult)

	// set status 201
	render.Status(r, http.StatusCreated)
	render.JSON(w, r, post)
}

func FindAllUsers(w http.ResponseWriter, r *http.Request) {
	rows, err := DB.Query("SELECT * FROM users")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	var users []User

	for rows.Next() {
		var user User

		err := rows.Scan(&user.ID, &user.Name, &user.Username, &user.Active)
		if err != nil {
			panic(err.Error())
		}

		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		panic(err.Error())
	}
	render.JSON(w, r, users)
}

func FindPostByID(w http.ResponseWriter, r *http.Request) {
	postID := chi.URLParam(r, "postID")

	var post Post

	row := DB.QueryRow("SELECT * FROM posts WHERE id = ?", postID)
	if err := row.Scan(&post.ID, &post.Title, &post.Body); err != nil {
		if err == sql.ErrNoRows {
			render.JSON(w, r, Error{Message: fmt.Sprintf("PostID %s not found", postID)})
			return
		}
		panic(err)
	}

	render.JSON(w, r, row)
}

func FindUserByID(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")

	var user User

	row := DB.QueryRow("SELECT * FROM users WHERE id = ?", userID)
	if err := row.Scan(&user.ID, &user.Name, &user.Username, &user.Active); err != nil {
		if err == sql.ErrNoRows {
			render.JSON(w, r, Error{Message: fmt.Sprintf("UserID %s not found", userID)})
			return
		}
		panic(err)
	}

	render.JSON(w, r, row)
}
