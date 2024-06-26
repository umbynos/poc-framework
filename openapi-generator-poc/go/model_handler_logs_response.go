// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Builder API Documentation
 *
 * This is the documentation for the Builder API service.
 *
 * API version: 2.0
 * Contact: support@swagger.io
 */

package openapi




// HandlerLogsResponse - LogsResponse represents a response for a compilation logs request
type HandlerLogsResponse struct {

	Stderr string `json:"stderr,omitempty"`

	Stdout string `json:"stdout,omitempty"`
}

// AssertHandlerLogsResponseRequired checks if the required fields are not zero-ed
func AssertHandlerLogsResponseRequired(obj HandlerLogsResponse) error {
	return nil
}

// AssertHandlerLogsResponseConstraints checks if the values respects the defined constraints
func AssertHandlerLogsResponseConstraints(obj HandlerLogsResponse) error {
	return nil
}
