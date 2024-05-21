package api

import (
	"encoding/base64"
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
func (s *CompilationsQueue) GetCompilationsId(w http.ResponseWriter, r *http.Request, id string) {
	if _, ok := s.Compilation[id]; !ok {
		w.WriteHeader(http.StatusNotFound)
		errStr := "Compilation not found"
		json.NewEncoder(w).Encode(HandlerErrNotFoundResponse{
			Err: &errStr,
		})
		return
	}
	status := s.Compilation[id].Status
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(HandlerCompilationResponse{
		Id:     &id,
		Status: &status,
	})
}

// Get the compilation arfitacts
// (GET /compilations/{id}/artifacts)
func (s *CompilationsQueue) GetCompilationsIdArtifacts(w http.ResponseWriter, r *http.Request, id string, params GetCompilationsIdArtifactsParams) {
	if _, ok := s.Compilation[id]; !ok {
		w.WriteHeader(http.StatusNotFound)
		errStr := "Compilation not found"
		json.NewEncoder(w).Encode(HandlerErrNotFoundResponse{
			Err: &errStr,
		})
		return
	}

	name := "test"
	binContent := base64.RawStdEncoding.EncodeToString([]byte("binary dat"))
	elfContent := base64.RawStdEncoding.EncodeToString([]byte("elf data"))
	hexContent := base64.RawStdEncoding.EncodeToString([]byte("hex data"))

	switch *params.Type {
	case "bin":
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(HandlerArtifactResponse{
			Name: &name,
			Bin:  &binContent,
		})
	case "elf":
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(HandlerArtifactResponse{
			Name: &name,
			Elf:  &elfContent,
		})
	case "hex":
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(HandlerArtifactResponse{
			Name: &name,
			Hex:  &hexContent,
		})
	default:
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(HandlerArtifactResponse{
			Name: &name,
			Bin:  &binContent,
			Elf:  &elfContent,
			Hex:  &hexContent,
		})
	}

}

// Stop a compilation
// (POST /compilations/{id}/cancel)
func (s *CompilationsQueue) PostCompilationsIdCancel(w http.ResponseWriter, r *http.Request, id string) {
	if _, ok := s.Compilation[id]; !ok {
		w.WriteHeader(http.StatusNotFound)
		errStr := "Compilation not found"
		json.NewEncoder(w).Encode(HandlerErrNotFoundResponse{
			Err: &errStr,
		})
		return
	}
	status := "cancelled"
	s.Compilation[id] = CompilationStatus{
		Compilation: s.Compilation[id].Compilation,
		Status:      status,
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(HandlerCompilationResponse{
		Id:     &id,
		Status: &status,
	})
}

// Get the compilation logs
// (GET /compilations/{id}/logs)
func (s *CompilationsQueue) GetCompilationsIdLogs(w http.ResponseWriter, r *http.Request, id string) {
	if _, ok := s.Compilation[id]; !ok {
		w.WriteHeader(http.StatusNotFound)
		errStr := "Compilation not found"
		json.NewEncoder(w).Encode(HandlerErrNotFoundResponse{
			Err: &errStr,
		})
		return
	}
	stdoutContent := "stdout"
	stderrContent := "stderr"
	json.NewEncoder(w).Encode(HandlerLogsResponse{
		Stdout: &stdoutContent,
		Stderr: &stderrContent,
	})
}
