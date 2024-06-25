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

func updateBookHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var updatedBook Book
	err := json.NewDecoder(r.Body).Decode(&updatedBook)
	if err != nil {
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": err.Error()})
		return
	}

	// Validate book data
	if updatedBook.Name == "" || updatedBook.Author == "" || updatedBook.Price <= 0 {
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Title, Author, and Price are required"})
		return
	}

	for index, book := range books {
		if book.ID == id {
			updatedBook.ID = book.ID
			books[index] = updatedBook
			w.Header().Set("content-type", "application/json")
			json.NewEncoder(w).Encode(updatedBook)
			return
		}
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"error": "Book not found!"})
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", homePageHandler).Methods("GET")
	router.HandleFunc("/books", getBooksHandler).Methods("GET")
	router.HandleFunc("/books", addBookHandler).Methods("POST")
	router.HandleFunc("/books/{id}", updateBookHandler).Methods("PUT")
	// Add handler to find a book by id and delete it from the list

	log.Println("Starting server on port 8081")
	http.ListenAndServe(":8081", router)
}
