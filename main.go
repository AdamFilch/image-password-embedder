package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// Create Temporary Directory for temp files
	os.MkdirAll("uploads", os.ModePerm)

	// Serve Static Files
	// http.Handle()

	// Initialize server
	port := ":8080"
	fmt.Println("Server running on http://localhost" + port)
	log.Fatal(http.ListenAndServe(port, nil))
}
