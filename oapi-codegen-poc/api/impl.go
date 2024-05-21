package api

import (
	"context"
	"encoding/base64"

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

var _ StrictServerInterface = (*CompilationsQueue)(nil)

// Get the status of the server
// (GET /alive)
func (s *CompilationsQueue) GetAlive(ctx context.Context, request GetAliveRequestObject) (GetAliveResponseObject, error) {
	result := "OK"
	return GetAlive200JSONResponse{Message: &result}, nil
}

// Add a new compilation
// (POST /compilations)
func (s *CompilationsQueue) PostCompilations(ctx context.Context, request PostCompilationsRequestObject) (PostCompilationsResponseObject, error) {
	compilation := request.Body
	uuid := uuid.New()
	statusCreated := "created"
	s.Compilation[uuid.String()] = CompilationStatus{
		Compilation: *compilation,
		Status:      statusCreated,
	}
	uuidStr := uuid.String()
	return PostCompilations201JSONResponse{Id: &uuidStr, Status: &statusCreated}, nil
}

// Get the status of a compilation
// (GET /compilations/{id})
func (s *CompilationsQueue) GetCompilationsId(ctx context.Context, request GetCompilationsIdRequestObject) (GetCompilationsIdResponseObject, error) {
	if _, ok := s.Compilation[request.Id]; !ok {
		errStr := "Compilation not found"
		return GetCompilationsId404JSONResponse{Err: &errStr}, nil
	}
	status := s.Compilation[request.Id].Status
	return GetCompilationsId200JSONResponse{Id: &request.Id, Status: &status}, nil
}

// Get the compilation arfitacts
// (GET /compilations/{id}/artifacts)
func (s *CompilationsQueue) GetCompilationsIdArtifacts(ctx context.Context, request GetCompilationsIdArtifactsRequestObject) (GetCompilationsIdArtifactsResponseObject, error) {
	if _, ok := s.Compilation[request.Id]; !ok {
		errStr := "Compilation not found"
		return GetCompilationsIdArtifacts404JSONResponse{Err: &errStr}, nil
	}

	name := "test"
	binContent := base64.RawStdEncoding.EncodeToString([]byte("binary dat"))
	elfContent := base64.RawStdEncoding.EncodeToString([]byte("elf data"))
	hexContent := base64.RawStdEncoding.EncodeToString([]byte("hex data"))
	binType := *request.Params.Type

	switch binType {
	case "bin":
		return GetCompilationsIdArtifacts200JSONResponse{
			Name: &name,
			Bin:  &binContent,
		}, nil
	case "elf":
		return GetCompilationsIdArtifacts200JSONResponse{
			Name: &name,
			Elf:  &elfContent,
		}, nil
	case "hex":
		return GetCompilationsIdArtifacts200JSONResponse{
			Name: &name,
			Hex:  &hexContent,
		}, nil
	default:
		return GetCompilationsIdArtifacts200JSONResponse{
			Name: &name,
			Bin:  &binContent,
			Elf:  &elfContent,
			Hex:  &hexContent,
		}, nil
	}
}

// Stop a compilation
// (POST /compilations/{id}/cancel)
func (s *CompilationsQueue) PostCompilationsIdCancel(ctx context.Context, request PostCompilationsIdCancelRequestObject) (PostCompilationsIdCancelResponseObject, error) {
	if _, ok := s.Compilation[request.Id]; !ok {
		errStr := "Compilation not found"
		return PostCompilationsIdCancel404JSONResponse{Err: &errStr}, nil
	}
	status := "cancelled"
	s.Compilation[request.Id] = CompilationStatus{
		Compilation: s.Compilation[request.Id].Compilation,
		Status:      status,
	}
	return PostCompilationsIdCancel200JSONResponse{Id: &request.Id, Status: &status}, nil
}

// Get the compilation logs
// (GET /compilations/{id}/logs)
func (s *CompilationsQueue) GetCompilationsIdLogs(ctx context.Context, request GetCompilationsIdLogsRequestObject) (GetCompilationsIdLogsResponseObject, error) {
	if _, ok := s.Compilation[request.Id]; !ok {
		errStr := "Compilation not found"
		return GetCompilationsIdLogs404JSONResponse{Err: &errStr}, nil
	}
	stdoutContent := "stdout"
	stderrContent := "stderr"
	return GetCompilationsIdLogs200JSONResponse{
		Stdout: &stdoutContent,
		Stderr: &stderrContent,
	}, nil
}
