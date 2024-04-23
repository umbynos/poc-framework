package main

import (
	"log"
	handler "net-http-poc/internal/v1"
	"net/http"
)

func main() {
	// Initialize the compilations queue
	handler.Init()
	// Define your API routes
	http.HandleFunc("GET /v1/alive", handler.Alive)
	http.HandleFunc("POST /v1/compilations", handler.CreateCompilation)
	http.HandleFunc("POST /v1/compilations/{id}/cancel", handler.StopCompilation)
	http.HandleFunc("GET /v1/compilations/{id}", handler.GetCompilations)
	http.HandleFunc("GET /v1/compilations/{id}/artifacts", handler.GetArtifacts)
	http.HandleFunc("GET /v1/compilations/{id}/logs", handler.GetLogs)

	// Start the server
	log.Fatal(http.ListenAndServe(":8080", nil))
}
