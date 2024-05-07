package handler

import (
	"encoding/json"
	"net/http"
)

// ErrorResponse represents an error response
// @Description ErrorResponse represents an error response
type ErrorResponse struct {
	Err string `json:"err" example:"error message"`
}

// AliveResponse represents a response for the alive endpoint
// @Description AliveResponse represents a response for the alive endpoint
type AliveResponse struct {
	Message string `json:"message" example:"OK"`
}

// CompilationResponse represents a response for a compilation request
// @Description CompilationResponse represents a response for a compilation request
type CompilationResponse struct {
	UUID   string `json:"id" example:"123e4567-e89b-12d3-a456-426614174000"`
	Status string `json:"status" example:"created"`
}

// ArtifactResponse represents a response for a compilation artifacts request
// @Description ArtifactResponse represents a response for a compilation artifacts request
type ArtifactResponse struct {
	Name string `json:"name" example:"test"`
	Bin  string `json:"bin" example:"YmluYXJ5IGRhdGE="`
	Elf  string `json:"elf" example:"ZWxmIGRhdGE="`
	Hex  string `json:"hex" example:"aGV4IGRhdGE="`
}

// LogsResponse represents a response for a compilation logs request
// @Description LogsResponse represents a response for a compilation logs request
type LogsResponse struct {
	Stdout string `json:"stdout" example:"Hello, World!"`
	Stderr string `json:"stderr" example:"Error: World not found"`
}

func encodeResponse(w http.ResponseWriter, statusCode int, resp interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if resp == nil {
		return
	}
	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		panic(err)
	}
}
