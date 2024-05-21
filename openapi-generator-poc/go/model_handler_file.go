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




// HandlerFile - File represents a file to be compiled
type HandlerFile struct {

	Data string `json:"data,omitempty"`

	Name string `json:"name,omitempty"`
}

// AssertHandlerFileRequired checks if the required fields are not zero-ed
func AssertHandlerFileRequired(obj HandlerFile) error {
	return nil
}

// AssertHandlerFileConstraints checks if the values respects the defined constraints
func AssertHandlerFileConstraints(obj HandlerFile) error {
	return nil
}
