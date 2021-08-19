package main

import (
	"fmt"
	"log"
	"net/http"
)

type store struct {
	StoreID   string `json:"id"`
	NameStore string `json:"name"`
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "pong")

	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/ping", pingHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
