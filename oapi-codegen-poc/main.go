package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"oapi-codegen-poc/api"
	"os"

	middleware "github.com/oapi-codegen/nethttp-middleware"
)

//go:embed static
var content embed.FS

func main() {

	swagger, err := api.GetSwagger()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading swagger spec\n: %s", err)
		os.Exit(1)
	}

	// Clear out the servers array in the swagger spec, that skips validating
	// that server names match. We don't know how this thing will be run.
	swagger.Servers = nil

	// create a type that satisfies the `api.ServerInterface`, which contains an implementation of every operation from the generated code
	server := api.NewCompilationQueue()

	r := http.NewServeMux()

	// get an `http.Handler` that we can use
	h := api.HandlerFromMux(server, r)

	fsys, _ := fs.Sub(content, "static")
	r.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.FS(fsys))))

	// Use our validation middleware to check all requests against the
	// OpenAPI schema.
	h = middleware.OapiRequestValidator(swagger)(h)

	s := &http.Server{
		Handler: h,
		Addr:    "0.0.0.0:8080",
	}

	// And we serve HTTP until the world ends.
	log.Fatal(s.ListenAndServe())
}
