package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
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

func main() {
	http.HandleFunc("/ping", pingHandler)

	http.HandleFunc("/stores", getAllStores)

	http.HandleFunc("/stores/", getSpecificStore)

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

var stores []Store

func getInformationsJSON(w http.ResponseWriter, r *http.Request) {

	byteValueJSON := readStoresFromArchive()

	err := json.Unmarshal(byteValueJSON, &stores)

	if err != nil {
		log.Fatal(err)
	}
}

func getAllStores(w http.ResponseWriter, r *http.Request) {

	if err := json.NewEncoder(w).Encode(stores); err != nil {
		log.Fatal(err)
	}
	// for _, store := range stores {
	// 	for i := range store.StoreEmployees {
	// 		_, err := fmt.Fprintf(
	// 			w, "StoreID: %v, StoreBrand: %v, StoreName: %v, StoreAddress: %v, %v, %v, Employee: [EmployeeID: %v, EmployeeName: %v]\n",
	// 			store.StoreID,
	// 			store.StoreBrand,
	// 			store.StoreName,
	// 			store.StoreAddress.City,
	// 			store.StoreAddress.State,
	// 			store.StoreAddress.Street,
	// 			store.StoreEmployees[i].EmployeeID,
	// 			store.StoreEmployees[i].EmployeeName,
	// 		)

	// 		if err != nil {
	// 			log.Fatal(err)
	// 		}
	// 	}
	// }
}

func getSpecificStore(w http.ResponseWriter, r *http.Request) {

	partsOfURL := strings.Split(r.URL.Path, "/")
	fmt.Println(partsOfURL)
	var brand string = partsOfURL[2]

	for _, store := range stores {
		if store.StoreBrand == brand {
			fmt.Fprintf(w, "%v\n", store)
		}
	}
}
