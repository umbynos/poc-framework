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




// HandlerLibrary - Library represents a library to be included in the compilation
type HandlerLibrary struct {

	Name string `json:"name,omitempty"`

	Version string `json:"version,omitempty"`
}

// AssertHandlerLibraryRequired checks if the required fields are not zero-ed
func AssertHandlerLibraryRequired(obj HandlerLibrary) error {
	return nil
}

// AssertHandlerLibraryConstraints checks if the values respects the defined constraints
func AssertHandlerLibraryConstraints(obj HandlerLibrary) error {
	return nil
}
