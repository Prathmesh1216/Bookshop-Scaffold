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

func getBooksHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func addBookHandler(w http.ResponseWriter, r *http.Request) {
	//	Add the implementation here
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", homePageHandler).Methods("GET")
	router.HandleFunc("/books", getBooksHandler).Methods("GET")
	// Add a new route to create a new book

	log.Println("Starting server on port 8081")
	http.ListenAndServe(":8081", router)
}
