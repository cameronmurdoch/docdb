package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type collection struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var collections = []collection{
	{ID: "2f1f71e5-e149-4cc7-b23a-3c3c3d39a4ee", Name: "default"},
	{ID: "6743c744-1a06-462e-82e6-85c9d0b2399f", Name: "Test"},
	{ID: "bc1f09ae-f9cc-45cd-b265-8b6950c5b714", Name: "Weatherdata"},
}

func Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	io.WriteString(w, `{"hello": "world"}`)
}

func Collections(w http.ResponseWriter, r *http.Request) {
	response, _ := json.Marshal(collections)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func main() {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", Home)
	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/collections", Collections)
	loggedRouter := handlers.LoggingHandler(os.Stdout, r)
	log.Fatal(http.ListenAndServe(":8080", loggedRouter))
}

// e6bbb1d6-2cfc-4271-b8f9-ed35b600f368
// 02b1890f-7668-4d1c-aa09-21586906131b
// 0fb890af-f2bf-44a7-8662-5771049e3cb6
