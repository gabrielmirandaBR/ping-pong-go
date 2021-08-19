package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Stores struct {
	Stores []Store `json:"acme-stores"`
}

type Store struct {
	StoreID        string         `json:"id"`
	StoreBrand     string         `json:"brand_label"`
	StoreName      string         `json:"name"`
	StoreAddress   StoreAddress   `json:"address"`
	StoreEmployees StoreEmployees `json:"employees"`
}

type StoreAddress struct {
	City   string `json:"city"`
	State  string `json:"state"`
	Street string `json:"street"`
}

type StoreEmployees struct {
	EmployeeID   string `json:"id"`
	EmployeeName string `json:"name"`
}

func main() {
	getInformationsJSON()

	http.HandleFunc("/ping", pingHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "pong")

	if err != nil {
		log.Fatal(err)
	}
}

func readStoresFromArchive() []byte {

	jsonFile, err := os.Open("acme-stores.json")

	if err != nil {
		fmt.Println("error", err)
	}

	fmt.Println("Successfully opened json file", jsonFile)

	byteValueJSON, err := ioutil.ReadAll(jsonFile)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully readed file in bytes", byteValueJSON)

	jsonFile.Close()

	return byteValueJSON
}

func getInformationsJSON() {
	var stores []Stores

	byteValueJSON := readStoresFromArchive()

	err := json.Unmarshal([]byte(byteValueJSON), &stores)

	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println(stores)
}
