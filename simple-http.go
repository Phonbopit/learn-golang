package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!")
	})

	http.HandleFunc("/json", handleJson)
	http.HandleFunc("/no-access", handleNoAccess)

	http.ListenAndServe(":7777", nil)
}

func handleJson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	res := make(map[string]string)
	res["message"] = "Hello World!"

	jsonResponse, err := json.Marshal(res)
	if err != nil {
		log.Fatalf("Error in JSON format. Err : %s", err)
	}

	w.Write(jsonResponse)
	return
}

func handleNoAccess(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusUnauthorized)
	w.Header().Set("Content-Type", "application/json")

	res := make(map[string]string)
	res["message"] = "Unauthorized"
	res["code"] = "10000"

	jsonResponse, err := json.Marshal(res)
	if err != nil {
		log.Fatalf("Error in JSON. Err : %s", err)
	}

	w.Write(jsonResponse)
}
