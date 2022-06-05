package main

import (
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	io.WriteString(w, `{"hello": "world"}`)
}

func Collections(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	io.WriteString(w, `{"hello": "world"}`)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Home)
	api := router.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/collections", Collections)
	log.Fatal(http.ListenAndServe(":8080", router))
}
