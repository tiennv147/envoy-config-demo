package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/api/v1/echo", echoHandler)
	http.HandleFunc("/api/v1/authz/", authzHandler)
	http.HandleFunc("/", everythingElseHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}
	fmt.Println("Request Echo:", r)
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func authzHandler(w http.ResponseWriter, r *http.Request) {
	// Print the request for debugging
	fmt.Println("Request AuthZ:", r)

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Authorized")
}

func everythingElseHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request All:", r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Everything else")
}
