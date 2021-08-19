package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func pingHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "pong")

	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	readStoresFromArchive()

	http.HandleFunc("/ping", pingHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func readStoresFromArchive() {

	jsonFile, err := os.Open("acme-stores.json")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	byteValueJSON, err := ioutil.ReadAll(jsonFile)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(byteValueJSON)

	defer jsonFile.Close()
}
