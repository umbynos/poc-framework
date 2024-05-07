package main

import (
	"log"
	handler "net-http-poc/internal/v1"
	"net/http"

	_ "net-http-poc/docs"

	httpSwagger "github.com/swaggo/http-swagger/v2"
)

//	@title			Builder API Documentation
//	@version		2.0
//	@description	This is the documentation for the Builder API service.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:8080
//	@BasePath	/v1

func main() {
	// Initialize the compilations queue
	handler.Init()
	// Define your API routes
	http.HandleFunc("GET /v1/docs/*", httpSwagger.Handler(httpSwagger.URL("/v1/docs/doc.json")))
	http.HandleFunc("GET /v1/alive", handler.Alive)
	http.HandleFunc("POST /v1/compilations", handler.CreateCompilation)
	http.HandleFunc("POST /v1/compilations/{id}/cancel", handler.StopCompilation)
	http.HandleFunc("GET /v1/compilations/{id}", handler.GetCompilations)
	http.HandleFunc("GET /v1/compilations/{id}/artifacts", handler.GetArtifacts)
	http.HandleFunc("GET /v1/compilations/{id}/logs", handler.GetLogs)

	// Start the server
	log.Fatal(http.ListenAndServe(":8080", nil))
}
