package main

import (
	"log"
	"net/http"
	openapi "openapi-generator-poc/go"
)

func main() {
	log.Printf("Server started")

	CompilationAPIService := openapi.NewCompilationAPIService()
	CompilationAPIController := openapi.NewCompilationAPIController(CompilationAPIService)

	StatusAPIService := openapi.NewStatusAPIService()
	StatusAPIController := openapi.NewStatusAPIController(StatusAPIService)

	router := openapi.NewRouter(CompilationAPIController, StatusAPIController)

	log.Fatal(http.ListenAndServe(":8080", router))
}
