package openapi

import (
	"context"
	"encoding/base64"

	"github.com/google/uuid"
)

type CompilationStatus struct {
	Compilation HandlerCompilation
	Status      string
}

// CompilationAPIService is a service that implements the logic for the CompilationAPIServicer
// This service should implement the business logic for every endpoint for the CompilationAPI API.
// Include any external packages or services that will be required by this service.
type CompilationAPIService struct {
	Compilation map[string]CompilationStatus
}

// NewCompilationAPIService creates a default api service
func NewCompilationAPIService() *CompilationAPIService {
	return &CompilationAPIService{
		Compilation: make(map[string]CompilationStatus),
	}
}

// CompilationsIdArtifactsGet - Get the compilation arfitacts
func (s *CompilationAPIService) CompilationsIdArtifactsGet(ctx context.Context, id string, type_ string) (ImplResponse, error) {
	if _, ok := s.Compilation[id]; !ok {
		// TODO: Uncomment the next line to return response Response(404, HandlerErrNotFoundResponse{}) or use other options such as http.Ok ...
		return Response(404, HandlerErrNotFoundResponse{
			Err: "Compilation not found",
		}), nil
	}
	name := "test"
	binContent := base64.RawStdEncoding.EncodeToString([]byte("binary dat"))
	elfContent := base64.RawStdEncoding.EncodeToString([]byte("elf data"))
	hexContent := base64.RawStdEncoding.EncodeToString([]byte("hex data"))

	switch type_ {
	case "bin":
		// TODO: Uncomment the next line to return response Response(200, HandlerArtifactResponse{}) or use other options such as http.Ok ...
		return Response(200, HandlerArtifactResponse{
			Name: name,
			Bin:  binContent,
		}), nil
	case "elf":
		// TODO: Uncomment the next line to return response Response(200, HandlerArtifactResponse{}) or use other options such as http.Ok ...
		return Response(200, HandlerArtifactResponse{
			Name: name,
			Elf:  elfContent,
		}), nil
	case "hex":
		// TODO: Uncomment the next line to return response Response(200, HandlerArtifactResponse{}) or use other options such as http.Ok ...
		return Response(200, HandlerArtifactResponse{
			Name: name,
			Hex:  hexContent,
		}), nil
	default:
		// TODO: Uncomment the next line to return response Response(200, HandlerArtifactResponse{}) or use other options such as http.Ok ...
		return Response(200, HandlerArtifactResponse{
			Name: name,
			Bin:  binContent,
			Elf:  elfContent,
			Hex:  hexContent,
		}), nil
	}

	// TODO - update CompilationsIdArtifactsGet with the required logic for this service method.
	// Add api_compilation_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(400, HandlerErrBadRequestResponse{}) or use other options such as http.Ok ...
	// return Response(400, HandlerErrBadRequestResponse{}), nil

	// TODO: Uncomment the next line to return response Response(401, HandlerErrUnauthorizedResponse{}) or use other options such as http.Ok ...
	// return Response(401, HandlerErrUnauthorizedResponse{}), nil

	// TODO: Uncomment the next line to return response Response(403, HandlerErrForbiddenResponse{}) or use other options such as http.Ok ...
	// return Response(403, HandlerErrForbiddenResponse{}), nil

	// TODO: Uncomment the next line to return response Response(405, HandlerErrMethodNotAllowedResponse{}) or use other options such as http.Ok ...
	// return Response(405, HandlerErrMethodNotAllowedResponse{}), nil

	// TODO: Uncomment the next line to return response Response(410, HandlerErrGoneResponse{}) or use other options such as http.Ok ...
	// return Response(410, HandlerErrGoneResponse{}), nil

	// TODO: Uncomment the next line to return response Response(429, HandlerErrRateLimitExceededResponse{}) or use other options such as http.Ok ...
	// return Response(429, HandlerErrRateLimitExceededResponse{}), nil

	// TODO: Uncomment the next line to return response Response(500, HandlerErrInternalServerErrorResponse{}) or use other options such as http.Ok ...
	// return Response(500, HandlerErrInternalServerErrorResponse{}), nil

	// TODO: Uncomment the next line to return response Response(503, HandlerErrServiceUnavailableResponse{}) or use other options such as http.Ok ...
	// return Response(503, HandlerErrServiceUnavailableResponse{}), nil

	// return Response(http.StatusNotImplemented, nil), errors.New("CompilationsIdArtifactsGet method not implemented")
}

// CompilationsIdCancelPost - Stop a compilation
func (s *CompilationAPIService) CompilationsIdCancelPost(ctx context.Context, id string) (ImplResponse, error) {
	if _, ok := s.Compilation[id]; !ok {
		// TODO: Uncomment the next line to return response Response(404, HandlerErrNotFoundResponse{}) or use other options such as http.Ok ...
		return Response(404, HandlerErrNotFoundResponse{
			Err: "Compilation not found",
		}), nil
	}
	status := "cancelled"
	s.Compilation[id] = CompilationStatus{
		Compilation: s.Compilation[id].Compilation,
		Status:      status,
	}
	// TODO: Uncomment the next line to return response Response(200, HandlerCompilationResponse{}) or use other options such as http.Ok ...
	return Response(200, HandlerCompilationResponse{
		Id:     id,
		Status: status,
	}), nil

	// TODO - update CompilationsIdCancelPost with the required logic for this service method.
	// Add api_compilation_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(400, HandlerErrBadRequestResponse{}) or use other options such as http.Ok ...
	// return Response(400, HandlerErrBadRequestResponse{}), nil

	// TODO: Uncomment the next line to return response Response(401, HandlerErrUnauthorizedResponse{}) or use other options such as http.Ok ...
	// return Response(401, HandlerErrUnauthorizedResponse{}), nil

	// TODO: Uncomment the next line to return response Response(403, HandlerErrForbiddenResponse{}) or use other options such as http.Ok ...
	// return Response(403, HandlerErrForbiddenResponse{}), nil

	// TODO: Uncomment the next line to return response Response(405, HandlerErrMethodNotAllowedResponse{}) or use other options such as http.Ok ...
	// return Response(405, HandlerErrMethodNotAllowedResponse{}), nil

	// TODO: Uncomment the next line to return response Response(410, HandlerErrGoneResponse{}) or use other options such as http.Ok ...
	// return Response(410, HandlerErrGoneResponse{}), nil

	// TODO: Uncomment the next line to return response Response(429, HandlerErrRateLimitExceededResponse{}) or use other options such as http.Ok ...
	// return Response(429, HandlerErrRateLimitExceededResponse{}), nil

	// TODO: Uncomment the next line to return response Response(500, HandlerErrInternalServerErrorResponse{}) or use other options such as http.Ok ...
	// return Response(500, HandlerErrInternalServerErrorResponse{}), nil

	// TODO: Uncomment the next line to return response Response(503, HandlerErrServiceUnavailableResponse{}) or use other options such as http.Ok ...
	// return Response(503, HandlerErrServiceUnavailableResponse{}), nil

	// return Response(http.StatusNotImplemented, nil), errors.New("CompilationsIdCancelPost method not implemented")
}

// CompilationsIdGet - Get the status of a compilation
func (s *CompilationAPIService) CompilationsIdGet(ctx context.Context, id string) (ImplResponse, error) {
	if _, ok := s.Compilation[id]; !ok {
		// TODO: Uncomment the next line to return response Response(404, HandlerErrNotFoundResponse{}) or use other options such as http.Ok ...
		return Response(404, HandlerErrNotFoundResponse{
			Err: "Compilation not found",
		}), nil
	}
	status := s.Compilation[id].Status
	// TODO: Uncomment the next line to return response Response(200, HandlerCompilationResponse{}) or use other options such as http.Ok ...
	return Response(200, HandlerCompilationResponse{
		Id:     id,
		Status: status,
	}), nil

	// TODO - update CompilationsIdGet with the required logic for this service method.
	// Add api_compilation_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(400, HandlerErrBadRequestResponse{}) or use other options such as http.Ok ...
	// return Response(400, HandlerErrBadRequestResponse{}), nil

	// TODO: Uncomment the next line to return response Response(401, HandlerErrUnauthorizedResponse{}) or use other options such as http.Ok ...
	// return Response(401, HandlerErrUnauthorizedResponse{}), nil

	// TODO: Uncomment the next line to return response Response(403, HandlerErrForbiddenResponse{}) or use other options such as http.Ok ...
	// return Response(403, HandlerErrForbiddenResponse{}), nil

	// TODO: Uncomment the next line to return response Response(405, HandlerErrMethodNotAllowedResponse{}) or use other options such as http.Ok ...
	// return Response(405, HandlerErrMethodNotAllowedResponse{}), nil

	// TODO: Uncomment the next line to return response Response(410, HandlerErrGoneResponse{}) or use other options such as http.Ok ...
	// return Response(410, HandlerErrGoneResponse{}), nil

	// TODO: Uncomment the next line to return response Response(429, HandlerErrRateLimitExceededResponse{}) or use other options such as http.Ok ...
	// return Response(429, HandlerErrRateLimitExceededResponse{}), nil

	// TODO: Uncomment the next line to return response Response(500, HandlerErrInternalServerErrorResponse{}) or use other options such as http.Ok ...
	// return Response(500, HandlerErrInternalServerErrorResponse{}), nil

	// TODO: Uncomment the next line to return response Response(503, HandlerErrServiceUnavailableResponse{}) or use other options such as http.Ok ...
	// return Response(503, HandlerErrServiceUnavailableResponse{}), nil

	// return Response(http.StatusNotImplemented, nil), errors.New("CompilationsIdGet method not implemented")
}

// CompilationsIdLogsGet - Get the compilation logs
func (s *CompilationAPIService) CompilationsIdLogsGet(ctx context.Context, id string) (ImplResponse, error) {
	if _, ok := s.Compilation[id]; !ok {
		// TODO: Uncomment the next line to return response Response(404, HandlerErrNotFoundResponse{}) or use other options such as http.Ok ...
		return Response(404, HandlerErrNotFoundResponse{
			Err: "Compilation not found",
		}), nil
	}
	// TODO: Uncomment the next line to return response Response(200, HandlerLogsResponse{}) or use other options such as http.Ok ...
	return Response(200, HandlerLogsResponse{
		Stdout: "stdout",
		Stderr: "stderr",
	}), nil

	// TODO - update CompilationsIdLogsGet with the required logic for this service method.
	// Add api_compilation_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(400, HandlerErrBadRequestResponse{}) or use other options such as http.Ok ...
	// return Response(400, HandlerErrBadRequestResponse{}), nil

	// TODO: Uncomment the next line to return response Response(401, HandlerErrUnauthorizedResponse{}) or use other options such as http.Ok ...
	// return Response(401, HandlerErrUnauthorizedResponse{}), nil

	// TODO: Uncomment the next line to return response Response(403, HandlerErrForbiddenResponse{}) or use other options such as http.Ok ...
	// return Response(403, HandlerErrForbiddenResponse{}), nil

	// TODO: Uncomment the next line to return response Response(405, HandlerErrMethodNotAllowedResponse{}) or use other options such as http.Ok ...
	// return Response(405, HandlerErrMethodNotAllowedResponse{}), nil

	// TODO: Uncomment the next line to return response Response(410, HandlerErrGoneResponse{}) or use other options such as http.Ok ...
	// return Response(410, HandlerErrGoneResponse{}), nil

	// TODO: Uncomment the next line to return response Response(429, HandlerErrRateLimitExceededResponse{}) or use other options such as http.Ok ...
	// return Response(429, HandlerErrRateLimitExceededResponse{}), nil

	// TODO: Uncomment the next line to return response Response(500, HandlerErrInternalServerErrorResponse{}) or use other options such as http.Ok ...
	// return Response(500, HandlerErrInternalServerErrorResponse{}), nil

	// TODO: Uncomment the next line to return response Response(503, HandlerErrServiceUnavailableResponse{}) or use other options such as http.Ok ...
	// return Response(503, HandlerErrServiceUnavailableResponse{}), nil

	// return Response(http.StatusNotImplemented, nil), errors.New("CompilationsIdLogsGet method not implemented")
}

// CompilationsPost - Add a new compilation
func (s *CompilationAPIService) CompilationsPost(ctx context.Context, value HandlerCompilation) (ImplResponse, error) {
	uuid := uuid.New()
	statusCreated := "created"
	s.Compilation[uuid.String()] = CompilationStatus{
		Compilation: value,
		Status:      statusCreated,
	}
	// TODO - update CompilationsPost with the required logic for this service method.
	// Add api_compilation_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(201, HandlerCompilationResponse{}) or use other options such as http.Ok ...
	return Response(201, HandlerCompilationResponse{
		Id:     uuid.String(),
		Status: statusCreated,
	}), nil

	// TODO: Uncomment the next line to return response Response(400, HandlerErrBadRequestResponse{}) or use other options such as http.Ok ...
	// return Response(400, HandlerErrBadRequestResponse{}), nil

	// TODO: Uncomment the next line to return response Response(401, HandlerErrUnauthorizedResponse{}) or use other options such as http.Ok ...
	// return Response(401, HandlerErrUnauthorizedResponse{}), nil

	// TODO: Uncomment the next line to return response Response(403, HandlerErrForbiddenResponse{}) or use other options such as http.Ok ...
	// return Response(403, HandlerErrForbiddenResponse{}), nil

	// TODO: Uncomment the next line to return response Response(405, HandlerErrMethodNotAllowedResponse{}) or use other options such as http.Ok ...
	// return Response(405, HandlerErrMethodNotAllowedResponse{}), nil

	// TODO: Uncomment the next line to return response Response(429, HandlerErrRateLimitExceededResponse{}) or use other options such as http.Ok ...
	// return Response(429, HandlerErrRateLimitExceededResponse{}), nil

	// TODO: Uncomment the next line to return response Response(500, HandlerErrInternalServerErrorResponse{}) or use other options such as http.Ok ...
	// return Response(500, HandlerErrInternalServerErrorResponse{}), nil

	// TODO: Uncomment the next line to return response Response(503, HandlerErrServiceUnavailableResponse{}) or use other options such as http.Ok ...
	// return Response(503, HandlerErrServiceUnavailableResponse{}), nil

	//return Response(http.StatusNotImplemented, nil), errors.New("CompilationsPost method not implemented")
}
