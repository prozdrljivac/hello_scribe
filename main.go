package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	const serverPort string = "8000"

	mux := http.NewServeMux()
	mux.HandleFunc("/", requestHandler)

	server := &http.Server{
		Addr:           ":" + serverPort,
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	fmt.Println("Server is running on port: ", serverPort)

	if err := server.ListenAndServe(); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}

func requestHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("This is the route of our API!")
}

// Learn why are these properties important
// ReadTimeout:
// WriteTimeout:
// MaxHeaderBytes:

// Create an API Point
// Create a docker image
