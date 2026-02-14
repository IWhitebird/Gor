// +build ignore

// Local development server for the Go API.
// Run with: go run dev-server.go
package main

import (
	"fmt"
	"log"
	"net/http"

	handler "gor-web/api"
)

func main() {
	http.HandleFunc("/api/run", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		handler.Handler(w, r)
	})

	fmt.Println("Go API server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
