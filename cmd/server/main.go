package main

import (
	"github.com/evgenzhu/go-metrics-app/internal/router"
	"log"
	"net/http"
)

func main() {
	mux := router.NewRouter()

	log.Println("Server is running on :8080")

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("server failed to start %v", err)
	}
}
