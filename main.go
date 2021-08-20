package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type Store struct {
	StoreID        string       `json:"id"`
	StoreBrand     string       `json:"brand_label"`
	StoreName      string       `json:"name"`
	StoreAddress   StoreAddress `json:"address"`
	StoreEmployees []Employee   `json:"employees"`
}

type StoreAddress struct {
	City   string `json:"city"`
	State  string `json:"state"`
	Street string `json:"street"`
}

type Employee struct {
	EmployeeID   string `json:"id"`
	EmployeeName string `json:"name"`
}

var stores []Store

func main() {
	getInformationsJSON()

	http.HandleFunc("/ping", pingHandler)

	http.HandleFunc("/stores", getAllStores)
	http.HandleFunc("/stores/", getSpecificStore)

	log.Fatal(http.ListenAndServe(":8081", nil))
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "pong")

	if err != nil {
		log.Fatal(err)
	}
}

func readStoresFromArchive() []byte {
	byteValueJSON, err := ioutil.ReadFile("./acme-stores.json")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully readed file in bytes")

	return byteValueJSON
}

func getInformationsJSON() {
	byteValueJSON := readStoresFromArchive()

	err := json.Unmarshal([]byte(byteValueJSON), &stores)

	if err != nil {
		log.Fatal(err)
	}
}

func getAllStores(w http.ResponseWriter, r *http.Request) {

	err := json.NewEncoder(w).Encode(stores)

	if err != nil {
		fmt.Fprint(w, err)
	}

}

func getSpecificStore(w http.ResponseWriter, r *http.Request) {

	partsOfURL := strings.Split(r.URL.Path, "/")

	brandStore := partsOfURL[2]

	if len(partsOfURL) > 3 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	var filteredStores []Store

	for _, store := range stores {
		if brandStore == store.StoreBrand {
			filteredStores = append(filteredStores, store)
		}
	}

	err := json.NewEncoder(w).Encode(filteredStores)

	if err != nil {
		log.Fatal(err)
	}
}
