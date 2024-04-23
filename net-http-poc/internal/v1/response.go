package handler

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Err string `json:"err" msgpack:"err"`
}

type AliveResponse struct {
	Message string `json:"message"`
}

type CompilationResponse struct {
	UUID   string `json:"id"`
	Status string `json:"status"`
}

type ArtifactResponse struct {
	Name string `json:"name"`
	Bin  string `json:"bin"`
	Elf  string `json:"elf"`
	Hex  string `json:"hex"`
}

type LogsResponse struct {
	Stdout string `json:"stdout"`
	Stderr string `json:"stderr"`
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
