package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/rs/xid"
	"log"
	"net/http"
)

type Book struct {
	ID              string  `json:"id"`
	Name            string  `json:"name"`
	Author          string  `json:"author"`
	PublicationYear string  `json:"publicationYear"`
	Price           float64 `json:"price"`
}

var books = []Book{
	{xid.New().String(), "Harry Potter", "J K Rowling", "2012", 200.0},
	{xid.New().String(), "Game of Life", "Sahana", "2008", 250.0},
}

func homePageHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Welcome to the Bookshop API!"})
}

// Handler to return all available books
func getBooksHandler(w http.ResponseWriter, r *http.Request) {

}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", homePageHandler).Methods("GET")
	// Add a new route to return all available books here.

	log.Println("Starting server on port 8081")
	http.ListenAndServe(":8081", router)
}
