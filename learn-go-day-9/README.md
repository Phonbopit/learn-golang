# Learn go day 9

Learn how to use [chi](https://go-chi.io)

## Installation

```
go get -u github.com/go-chi/chi/v5
```

## API

### GET /posts

```
curl -X GET http://localhost:9998/posts
```

### GET /posts/{id}

```
curl -X GET http://localhost:9998/posts/1
```

### POST /posts

```
curl -X POST http://localhost:9998/posts \
-H "Content-Type: application/json" \
-d '{
    "title": "My first post",
    "body": "This is my first post",
    "author_id": 2
}'
```
