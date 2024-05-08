package handler

import (
	"encoding/json"
	"net/http"
)

// ErrBadRequestResponse represents a bad request response
// @Description ErrBadRequestResponse represents a bad request response
type ErrBadRequestResponse struct {
	Err string `json:"err" example:"bad request"`
}

// ErrUnauthorizedResponse represents an unauthorized response
// @Description ErrUnauthorizedResponse represents an unauthorized response
type ErrUnauthorizedResponse struct {
	Err string `json:"err" example:"unauthorized"`
}

// ErrForbiddenResponse represents a forbidden response
// @Description ErrForbiddenResponse represents a forbidden response
type ErrForbiddenResponse struct {
	Err string `json:"err" example:"forbidden"`
}

// ErrNotFoundResponse represents a not found response
// @Description ErrNotFoundResponse represents a not found response
type ErrNotFoundResponse struct {
	Err string `json:"err" example:"not found"`
}

// ErrMethodNotAllowedResponse represents a method not allowed response
// @Description ErrMethodNotAllowedResponse represents a method not allowed response
type ErrMethodNotAllowedResponse struct {
	Err string `json:"err" example:"method not allowed"`
}

// ErrGoneResponse represents a gone response
// @Description ErrGoneResponse represents a gone response
type ErrGoneResponse struct {
	Err string `json:"err" example:"gone"`
}

// ErrRateLimitExceededResponse represents a rate limit exceeded response
// @Description ErrRateLimitExceededResponse represents a rate limit exceeded response
type ErrRateLimitExceededResponse struct {
	Err string `json:"err" example:"rate limit exceeded"`
}

// ErrInternalServerErrorResponse represents an internal server error response
// @Description ErrInternalServerErrorResponse represents an internal server error response
type ErrInternalServerErrorResponse struct {
	Err string `json:"err" example:"internal server error"`
}

// ErrServiceUnavailableResponse represents a service unavailable response
// @Description ErrServiceUnavailableResponse represents a service unavailable response
type ErrServiceUnavailableResponse struct {
	Err string `json:"err" example:"service unavailable"`
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
