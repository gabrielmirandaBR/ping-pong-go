package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

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
	Employees []Employee `json:"employees"`
}

type Employee struct {
	EmployeeID   string `json:"id"`
	EmployeeName string `json:"name"`
}

func main() {
	http.HandleFunc("/ping", pingHandler)
	http.HandleFunc("/stores", getInformationsJSON)

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

	fmt.Println("Successfully opened json file")

	byteValueJSON, err := ioutil.ReadAll(jsonFile)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully readed file in bytes")

	jsonFile.Close()

	return byteValueJSON
}

func getInformationsJSON(w http.ResponseWriter, r *http.Request) {
	var stores []Store

	byteValueJSON := readStoresFromArchive()

	err := json.Unmarshal([]byte(byteValueJSON), &stores)

	if err != nil {
		fmt.Println("error:", err)
	}

	for i := 0; i < len(stores); i += 1 {
		fmt.Fprintf(
			w, "StoreID: %v, StoreBrand: %v, StoreName: %v, StoreAddress: %v, %v, %v\n",
			stores[i].StoreID,
			stores[i].StoreBrand,
			stores[i].StoreName,
			stores[i].StoreAddress.City,
			stores[i].StoreAddress.State,
			stores[i].StoreAddress.Street,
		)
	}
}
