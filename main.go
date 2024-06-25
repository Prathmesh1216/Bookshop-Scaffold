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
	var newBook Book
	json.NewDecoder(r.Body).Decode(&newBook)

	// Validate book data
	if newBook.Name == "" || newBook.Author == "" || newBook.Price <= 0 {
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Title, Author, and Price are required"})
		return
	}

	newBook.ID = xid.New().String()
	books = append(books, newBook)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newBook)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", homePageHandler).Methods("GET")
	router.HandleFunc("/books", getBooksHandler).Methods("GET")
	router.HandleFunc("/books", addBookHandler).Methods("POST")
	// Add handler to find a book by id and update it with the given input params

	log.Println("Starting server on port 8081")
	http.ListenAndServe(":8081", router)
}
