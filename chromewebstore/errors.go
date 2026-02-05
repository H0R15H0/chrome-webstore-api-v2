package chromewebstore

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// googleAPIErrorResponse represents the error response structure from Google APIs.
type googleAPIErrorResponse struct {
	Error *googleAPIErrorDetail `json:"error"`
}

// googleAPIErrorDetail represents the error detail structure.
type googleAPIErrorDetail struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// APIError represents an error returned by the Chrome Web Store API.
type APIError struct {
	// StatusCode is the HTTP status code.
	StatusCode int
	// Status is the HTTP status text.
	Status string
	// Body is the response body.
	Body string
	// Message provides a human-readable error message.
	Message string
}

// Error returns the error message.
func (e *APIError) Error() string {
	if e.Message != "" {
		return fmt.Sprintf("chromewebstore: %s (HTTP %d)", e.Message, e.StatusCode)
	}
	return fmt.Sprintf("chromewebstore: HTTP %d: %s", e.StatusCode, e.Status)
}

// IsNotFound returns true if the error is a 404 Not Found error.
func (e *APIError) IsNotFound() bool {
	return e.StatusCode == http.StatusNotFound
}

// IsUnauthorized returns true if the error is a 401 Unauthorized error.
func (e *APIError) IsUnauthorized() bool {
	return e.StatusCode == http.StatusUnauthorized
}

// IsForbidden returns true if the error is a 403 Forbidden error.
func (e *APIError) IsForbidden() bool {
	return e.StatusCode == http.StatusForbidden
}

// IsRateLimited returns true if the error is a 429 Too Many Requests error.
func (e *APIError) IsRateLimited() bool {
	return e.StatusCode == http.StatusTooManyRequests
}

// newAPIError creates a new APIError from an HTTP response.
func newAPIError(resp *http.Response, body []byte) *APIError {
	apiErr := &APIError{
		StatusCode: resp.StatusCode,
		Status:     resp.Status,
		Body:       string(body),
	}

	// Try to parse the error response as Google API error format
	var errResp googleAPIErrorResponse
	if err := json.Unmarshal(body, &errResp); err == nil && errResp.Error != nil {
		apiErr.Message = errResp.Error.Message
	}

	return apiErr
}
