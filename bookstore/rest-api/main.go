package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"math/rand"
	"net/http"
	"strconv"
)

type BookStore struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

var bookstores []BookStore

func getbooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bookstores)
}

func createbook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var post BookStore
	_ = json.NewDecoder(r.Body).Decode(&post)
	post.ID = strconv.Itoa(rand.Intn(1000000))
	bookstores = append(bookstores, post)
	json.NewEncoder(w).Encode(&post)
}

func getbook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range bookstores {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&BookStore{})
}

func updatebook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range bookstores {
		if item.ID == params["id"] {
			bookstores = append(bookstores[:index], bookstores[index+1:]...)
			var post BookStore
			_ = json.NewDecoder(r.Body).Decode(&post)
			post.ID = params["id"]
			bookstores = append(bookstores, post)
			json.NewEncoder(w).Encode(&post)
			return
		}
	}
	json.NewEncoder(w).Encode(bookstores)
}

func deletebook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range bookstores {
		if item.ID == params["id"] {
			bookstores = append(bookstores[:index], bookstores[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(bookstores)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Everyone!"))
}

func main() {
	router := mux.NewRouter()
	bookstores = append(bookstores, BookStore{ID: "1", Name: "My first book", Description: "This is the content of my first book"})
	bookstores = append(bookstores, BookStore{ID: "2", Name: "My second book", Description: "This is the content of my second book"})

	router.HandleFunc("/books", getbooks).Methods("GET")
	router.HandleFunc("/books", createbook).Methods("POST")
	router.HandleFunc("/books/{id}", getbook).Methods("GET")
	router.HandleFunc("/books/{id}", updatebook).Methods("PUT")
	router.HandleFunc("/books/{id}", deletebook).Methods("DELETE")
	router.HandleFunc("/", handler)
	http.ListenAndServe(":8000", router)
}
