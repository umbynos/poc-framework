package openapi

import (
	"context"
)

// StatusAPIService is a service that implements the logic for the StatusAPIServicer
// This service should implement the business logic for every endpoint for the StatusAPI API.
// Include any external packages or services that will be required by this service.
type StatusAPIService struct {
}

// NewStatusAPIService creates a default api service
func NewStatusAPIService() *StatusAPIService {
	return &StatusAPIService{}
}

// AliveGet - Get the status of the server
func (s *StatusAPIService) AliveGet(ctx context.Context) (ImplResponse, error) {
	// TODO - update AliveGet with the required logic for this service method.
	// Add api_status_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, HandlerAliveResponse{}) or use other options such as http.Ok ...
	return Response(200, HandlerAliveResponse{"OK"}), nil

	// TODO: Uncomment the next line to return response Response(400, HandlerErrBadRequestResponse{}) or use other options such as http.Ok ...
	// return Response(400, HandlerErrBadRequestResponse{}), nil

	// TODO: Uncomment the next line to return response Response(405, HandlerErrMethodNotAllowedResponse{}) or use other options such as http.Ok ...
	// return Response(405, HandlerErrMethodNotAllowedResponse{}), nil

	// TODO: Uncomment the next line to return response Response(500, HandlerErrInternalServerErrorResponse{}) or use other options such as http.Ok ...
	// return Response(500, HandlerErrInternalServerErrorResponse{}), nil

	// TODO: Uncomment the next line to return response Response(503, HandlerErrServiceUnavailableResponse{}) or use other options such as http.Ok ...
	// return Response(503, HandlerErrServiceUnavailableResponse{}), nil

	// return Response(http.StatusNotImplemented, nil), errors.New("AliveGet method not implemented")
}
