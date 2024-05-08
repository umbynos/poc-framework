package handler

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
)

// Alive returns a simple message to indicate that the server is alive
// @Summary		Get the status of the server
// @Description	Get the status of the server
// @Tags		status
// @Produce		json
// @Success		200		{object}	AliveResponse
// @Failure		400		{object}	ErrBadRequestResponse
// @Failure		405		{object}	ErrMethodNotAllowedResponse
// @Failure		500 	{object}	ErrInternalServerErrorResponse
// @Failure		503 	{object}	ErrServiceUnavailableResponse
// @Router		/alive [get]
func Alive(w http.ResponseWriter, r *http.Request) {
	log.Println("I'm ALIVEEEEEEE")
	encodeResponse(w, http.StatusOK, AliveResponse{
		Message: "OK",
	})
}

// CreateCompilation adds a compilation request to the queue
// Parse the request body to get the compilation data
// Generate a new UUID for the compilation
// Add the compilation to the compilations queue
// Return the created compilation as JSON
// @Summary		Add a new compilation
// @Description	Add a new compilation to the queue
// @Tags		compilation
// @Accept		json
// @Produce		json
// @Param		value	body		Compilation	true	"Compilation object that needs to be added to the queue"
// @Success		201		{object}	CompilationResponse
// @Failure		400		{object}	ErrBadRequestResponse
// @Failure		401 	{object}	ErrUnauthorizedResponse
// @Failure		403 	{object}	ErrForbiddenResponse
// @Failure		405 	{object}	ErrMethodNotAllowedResponse
// @Failure		429		{object}	ErrRateLimitExceededResponse
// @Failure		500		{object}	ErrInternalServerErrorResponse
// @Failure		503 	{object}	ErrServiceUnavailableResponse
// @Router		/compilations [post]
func CreateCompilation(w http.ResponseWriter, r *http.Request) {
	var compilation Compilation
	err := json.NewDecoder(r.Body).Decode(&compilation)
	if err != nil {
		encodeResponse(w, http.StatusBadRequest, ErrBadRequestResponse{
			Err: err.Error(),
		})
		return
	}
	uuid := uuid.New()
	compilationsQueue[uuid.String()] = CompilationStatus{
		Compilation: compilation,
		Status:      "created",
	}
	log.Println("Compilation created:", uuid.String())
	encodeResponse(w, http.StatusCreated, CompilationResponse{
		UUID:   uuid.String(),
		Status: compilationsQueue[uuid.String()].Status,
	})
}

// Stop an existing compilation
// Extract the ID from the request URL
// Find the compilation in the compilations queue
// Update the compilation status to "cancelled"
// Return a success message
// @Summary		Stop a compilation
// @Description	Stop a compilation in the queue
// @Tags		compilation
// @Produce		json
// @Param		id		path 		string	true	"Compilation object that needs to be stopped"
// @Success		200		{object}	CompilationResponse
// @Failure		400		{object}	ErrBadRequestResponse
// @Failure		401	 	{object}	ErrUnauthorizedResponse
// @Failure		403		{object}	ErrForbiddenResponse
// @Failure		404		{object}	ErrNotFoundResponse
// @Failure		405  	{object}	ErrMethodNotAllowedResponse
// @Failure		410		{object}	ErrGoneResponse
// @Failure		429		{object}	ErrRateLimitExceededResponse
// @Failure		500		{object}	ErrInternalServerErrorResponse
// @Failure		503		{object}	ErrServiceUnavailableResponse
// @Router		/compilations/{id}/cancel [post]
func StopCompilation(w http.ResponseWriter, r *http.Request) {
	uuid := r.PathValue("id")
	if _, ok := compilationsQueue[uuid]; !ok {
		encodeResponse(w, http.StatusNotFound, ErrNotFoundResponse{
			Err: "Compilation not found",
		})
		return
	}
	compilationsQueue[uuid] = CompilationStatus{
		Compilation: compilationsQueue[uuid].Compilation,
		Status:      "cancelled",
	}
	log.Println("Compilation cancelled:", uuid)
	encodeResponse(w, http.StatusOK, CompilationResponse{
		UUID:   uuid,
		Status: compilationsQueue[uuid].Status,
	})
	log.Println("Compilation cancelled:", uuid)
}

// Extract the ID from the request URL
// Find the compilation in the compilations queue
// Return a success message
// @Summary		Get the status of a compilation
// @Description	Get the status of a compilation
// @Tags		compilation
// @Produce		json
// @Param		id		path 		string	true	"Compilation object we want to know the status of"
// @Success		200		{object}	CompilationResponse
// @Failure		400		{object}	ErrBadRequestResponse
// @Failure		401		{object}	ErrUnauthorizedResponse
// @Failure		403		{object}	ErrForbiddenResponse
// @Failure		404		{object}	ErrNotFoundResponse
// @Failure		405		{object}	ErrMethodNotAllowedResponse
// @Failure		410		{object}	ErrGoneResponse
// @Failure		429		{object}	ErrRateLimitExceededResponse
// @Failure		500		{object}	ErrInternalServerErrorResponse
// @Failure		503		{object}	ErrServiceUnavailableResponse
// @Router		/compilations/{id} [get]
func GetCompilations(w http.ResponseWriter, r *http.Request) {
	uuid := r.PathValue("id")
	if _, ok := compilationsQueue[uuid]; !ok {
		encodeResponse(w, http.StatusNotFound, ErrNotFoundResponse{
			Err: "Compilation not found",
		})
		return
	}
	log.Println("Compilation found:", uuid)
	encodeResponse(w, http.StatusOK, CompilationResponse{
		UUID:   uuid,
		Status: compilationsQueue[uuid].Status,
	})
}

// Get compilation artifacts
// Extract the ID from the request URL
// Find the compilation in the compilations slice
// Return the compilation artifacts as JSON
// @Summary		Get the compilation arfitacts
// @Description	Get the compilation artifacts
// @Tags		compilation
// @Produce		json
// @Param		id		path 		string	true	"Compilation object we want to retrieve the artifacts of"
// @Param		type	query		string	false	"Type of artifact we want to retrieve: bin, elf, hex"
// @Success		200		{object}	ArtifactResponse
// @Failure		400		{object}	ErrBadRequestResponse
// @Failure		401		{object}	ErrUnauthorizedResponse
// @Failure		403		{object}	ErrForbiddenResponse
// @Failure		404		{object}	ErrNotFoundResponse
// @Failure		405		{object}	ErrMethodNotAllowedResponse
// @Failure		410		{object}	ErrGoneResponse
// @Failure		429		{object}	ErrRateLimitExceededResponse
// @Failure		500		{object}	ErrInternalServerErrorResponse
// @Failure		503		{object}	ErrServiceUnavailableResponse
// @Router		/compilations/{id}/artifacts [get]
func GetArtifacts(w http.ResponseWriter, r *http.Request) {
	uuid := r.PathValue("id")
	if _, ok := compilationsQueue[uuid]; !ok {
		encodeResponse(w, http.StatusNotFound, ErrNotFoundResponse{
			Err: "Compilation not found",
		})
		return
	}
	log.Println("Compilation artifacts:", uuid)
	switch binType := r.URL.Query().Get("type"); binType {
	case "bin":
		encodeResponse(w, http.StatusOK, ArtifactResponse{
			Name: "test",
			Bin:  base64.RawStdEncoding.EncodeToString([]byte("binary data")),
		})
		return
	case "elf":
		encodeResponse(w, http.StatusOK, ArtifactResponse{
			Name: "test",
			Elf:  base64.RawStdEncoding.EncodeToString([]byte("elf data")),
		})
		return
	case "hex":
		encodeResponse(w, http.StatusOK, ArtifactResponse{
			Name: "test",
			Hex:  base64.RawStdEncoding.EncodeToString([]byte("hex data")),
		})
		return
	default:
		encodeResponse(w, http.StatusOK, ArtifactResponse{
			Name: "test",
			Bin:  base64.RawStdEncoding.EncodeToString([]byte("binary data")),
			Elf:  base64.RawStdEncoding.EncodeToString([]byte("elf data")),
			Hex:  base64.RawStdEncoding.EncodeToString([]byte("hex data")),
		})
		return
	}

}

// Get compilation logs
// Extract the ID from the request URL
// Find the compilation in the compilations slice
// Return the compilation logs as JSON
// @Summary		Get the compilation logs
// @Description	Get the compilation logs
// @Tags		compilation
// @Produce		json
// @Param		id		path 		string	true	"Compilation object we want to retrieve the logs of"
// @Success		200		{object}	LogsResponse
// @Failure		400		{object}	ErrBadRequestResponse
// @Failure		401		{object}	ErrUnauthorizedResponse
// @Failure		403 	{object}	ErrForbiddenResponse
// @Failure		404		{object}	ErrNotFoundResponse
// @Failure		405		{object}	ErrMethodNotAllowedResponse
// @Failure		410		{object}	ErrGoneResponse
// @Failure		429		{object}	ErrRateLimitExceededResponse
// @Failure		500		{object}	ErrInternalServerErrorResponse
// @Failure		503		{object}	ErrServiceUnavailableResponse
// @Router		/compilations/{id}/logs [get]
func GetLogs(w http.ResponseWriter, r *http.Request) {
	uuid := r.PathValue("id")
	if _, ok := compilationsQueue[uuid]; !ok {
		encodeResponse(w, http.StatusNotFound, ErrNotFoundResponse{
			Err: "Compilation not found",
		})
		return
	}
	log.Println("Compilation logs:", uuid)
	encodeResponse(w, http.StatusOK, LogsResponse{
		Stdout: "stdout",
		Stderr: "stderr",
	})
}
