package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/ping", pingHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
