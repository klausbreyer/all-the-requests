package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	separator := strings.Repeat("-", 50)
	fmt.Println(separator)
	fmt.Printf("Timestamp: %s\n", time.Now().Format(time.RFC3339))
	fmt.Printf("Method: %s\n", r.Method)
	fmt.Printf("Path: %s\n", r.URL.Path)

	fmt.Println("Headers:")
	for key, values := range r.Header {
		for _, value := range values {
			fmt.Printf("  %s: %s\n", key, value)
		}
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read body", http.StatusBadRequest)
		return
	}

	fmt.Println(separator)
	fmt.Printf("Payload: %s\n", body)
	fmt.Println(separator)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Received"))
}

func main() {
	port := flag.String("port", "8080", "Specify the port to listen on")
	flag.Parse()

	http.HandleFunc("/", handler)
	fmt.Printf("Starting server on :%s\n", *port)
	fmt.Println("Server is ready to handle requests and display all details.")

	if err := http.ListenAndServe(":"+*port, nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
