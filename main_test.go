package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

const ContentType = "Content-Type"
const ApplicationJson = "application/json"

func TestHome(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Home)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong statue code: got %v want %v",
			status, http.StatusOK)
	}

	if ctype := rr.Header().Get(ContentType); ctype != ApplicationJson {
		t.Errorf("content type header does not match: got %v want %v",
			ctype, "application/json")
	}

	expected := `{"hello": "world"}`
	if rr.Body.String() != expected {
		t.Errorf("hander returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestCollections(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/v1/collections", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Collections)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong statue code: got %v want %v",
			status, http.StatusOK)
	}

	if ctype := rr.Header().Get(ContentType); ctype != ApplicationJson {
		t.Errorf("content type header does not match: got %v want %v",
			ctype, "application/json")
	}
	/*
	   expected := `{"hello": "world"}`
	   if rr.Body.String() != expected {
	   	t.Errorf("hander returned unexpected body: got %v want %v",
	   		rr.Body.String(), expected)
	   } */
}
