package entities

// ErrorResponse defines the structure of the error response returned by the API.
type ErrorResponse struct {
	Message string `json:"message"` // Descriptive error message
	Code    int    `json:"code"`    // HTTP status code
}
