package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println()

	// Add current timestamp before output
	fmt.Printf("Timestamp: %s\n", time.Now().Format(time.RFC3339))
	// Log the request path
	fmt.Printf("Path: %s\n", r.URL.Path)

	// Log the request headers
	fmt.Println("Headers:")
	for key, values := range r.Header {
		for _, value := range values {
			fmt.Printf("%s: %s\n", key, value)
		}
	}

	// Read and log the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read body", http.StatusBadRequest)
		return
	}

	// Log the request payload
	fmt.Printf("Payload: %s\n", body)

	// Respond with a simple message
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Received"))
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
