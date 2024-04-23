package handler

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
)

func Alive(w http.ResponseWriter, r *http.Request) {
	log.Println("I'm ALIVEEEEEEE")
	encodeResponse(w, http.StatusOK, AliveResponse{
		Message: "OK",
	})
}

// Create a new compilation
// Parse the request body to get the compilation data
// Generate a new ID for the compilation
// Add the compilation to the compilations queue
// Return the created compilation as JSON
func CreateCompilation(w http.ResponseWriter, r *http.Request) {
	var compilation Compilation
	err := json.NewDecoder(r.Body).Decode(&compilation)
	if err != nil {
		encodeResponse(w, http.StatusBadRequest, ErrorResponse{
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
func StopCompilation(w http.ResponseWriter, r *http.Request) {
	uuid := r.PathValue("id")
	if _, ok := compilationsQueue[uuid]; !ok {
		encodeResponse(w, http.StatusNotFound, ErrorResponse{
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
func GetCompilations(w http.ResponseWriter, r *http.Request) {
	uuid := r.PathValue("id")
	if _, ok := compilationsQueue[uuid]; !ok {
		encodeResponse(w, http.StatusNotFound, ErrorResponse{
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
func GetArtifacts(w http.ResponseWriter, r *http.Request) {
	uuid := r.PathValue("id")
	if _, ok := compilationsQueue[uuid]; !ok {
		encodeResponse(w, http.StatusNotFound, ErrorResponse{
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
func GetLogs(w http.ResponseWriter, r *http.Request) {
	uuid := r.PathValue("id")
	if _, ok := compilationsQueue[uuid]; !ok {
		encodeResponse(w, http.StatusNotFound, ErrorResponse{
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
