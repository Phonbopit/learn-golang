package main

import (
  "fmt"
  "os"

)

import _ "github.com/joho/godotenv/autoload"

func main() {
  environment := os.Getenv("GO_ENV")
  secretKey := os.Getenv("SECRET_KEY")
  fmt.Println("GO_ENV: ", environment)
  fmt.Println("SECRET_KEY: ", secretKey)

  gogetenv()
}

func gogetenv() {
  fmt.Println("============\ngogetenv()\n============")
  // err := godotenv.Load()
  // if err != nil {
  //   log.Fatal("Error loading env")
  // }
  environment := os.Getenv("GO_ENV")
  secretKey := os.Getenv("SECRET_KEY")

  fmt.Println("GO_ENV: ", environment)
  fmt.Println("SECRET_KEY: ", secretKey)
}

