package api

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

type CompilationsQueue struct {
	Compilation map[string]CompilationStatus
}

// CompilationStatus represents the status of a compilation
type CompilationStatus struct {
	Compilation HandlerCompilation `json:"compilation"`
	Status      string             `json:"status"`
}

// NewCompilationQueue creates a new CompilationQueue
func NewCompilationQueue() *CompilationsQueue {
	return &CompilationsQueue{
		Compilation: make(map[string]CompilationStatus),
	}
}

var _ ServerInterface = (*CompilationsQueue)(nil)

// Get the status of the server
// (GET /alive)
func (s *CompilationsQueue) GetAlive(w http.ResponseWriter, r *http.Request) {
	result := "OK"
	resp := HandlerAliveResponse{
		Message: &result,
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

// Add a new compilation
// (POST /compilations)
func (s *CompilationsQueue) PostCompilations(w http.ResponseWriter, r *http.Request) {
	var compilation HandlerCompilation
	err := json.NewDecoder(r.Body).Decode(&compilation)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errStr := err.Error()
		json.NewEncoder(w).Encode(HandlerErrBadRequestResponse{
			Err: &errStr,
		})
		return
	}
	uuid := uuid.New()
	statusCreated := "created"
	s.Compilation[uuid.String()] = CompilationStatus{
		Compilation: compilation,
		Status:      statusCreated,
	}
	w.WriteHeader(http.StatusCreated)
	uuidStr := uuid.String()
	json.NewEncoder(w).Encode(HandlerCompilationResponse{
		Id:     &uuidStr,
		Status: &statusCreated,
	})
}

// Get the status of a compilation
// (GET /compilations/{id})
func (s *CompilationsQueue) GetCompilationsId(w http.ResponseWriter, r *http.Request, id string) {}

// Get the compilation arfitacts
// (GET /compilations/{id}/artifacts)
func (s *CompilationsQueue) GetCompilationsIdArtifacts(w http.ResponseWriter, r *http.Request, id string, params GetCompilationsIdArtifactsParams) {
}

// Stop a compilation
// (POST /compilations/{id}/cancel)
func (s *CompilationsQueue) PostCompilationsIdCancel(w http.ResponseWriter, r *http.Request, id string) {
}

// Get the compilation logs
// (GET /compilations/{id}/logs)
func (s *CompilationsQueue) GetCompilationsIdLogs(w http.ResponseWriter, r *http.Request, id string) {
}