package main

import (
	"fmt"
	"log"
	"net/http"

	ourcode "main.go/handlers"
)


func main() {
	http.HandleFunc("/", ourcode.HomeHandler)
	http.HandleFunc("/ascii-art", ourcode.AsciiArtHandler)
	fmt.Println("Server starting on :8080")
	fmt.Println("Visit: http://localhost:8080")
	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		fmt.Printf("Server failed to start: %v\n", err)
		log.Fatal(err)
	}
}
