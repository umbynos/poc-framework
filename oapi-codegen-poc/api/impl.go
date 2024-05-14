package api

import (
	"encoding/json"
	"net/http"
)

type Server struct{}

func NewServer() Server {
	return Server{}
}

// Get the status of the server
// (GET /alive)
func (s Server) GetAlive(w http.ResponseWriter, r *http.Request) {
	resp := HandlerAliveResponse{
		//Message: "OK",
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

// Add a new compilation
// (POST /compilations)
func (s Server) PostCompilations(w http.ResponseWriter, r *http.Request) {

}

// Get the status of a compilation
// (GET /compilations/{id})
func (s Server) GetCompilationsId(w http.ResponseWriter, r *http.Request, id string) {}

// Get the compilation arfitacts
// (GET /compilations/{id}/artifacts)
func (s Server) GetCompilationsIdArtifacts(w http.ResponseWriter, r *http.Request, id string, params GetCompilationsIdArtifactsParams) {
}

// Stop a compilation
// (POST /compilations/{id}/cancel)
func (s Server) PostCompilationsIdCancel(w http.ResponseWriter, r *http.Request, id string) {}

// Get the compilation logs
// (GET /compilations/{id}/logs)
func (s Server) GetCompilationsIdLogs(w http.ResponseWriter, r *http.Request, id string) {}
