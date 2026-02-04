package chromewebstore

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const (
	// DefaultBaseURL is the default base URL for the Chrome Web Store API.
	DefaultBaseURL = "https://chromewebstore.googleapis.com"
	// DefaultUploadBaseURL is the default base URL for upload operations.
	DefaultUploadBaseURL = "https://chromewebstore.googleapis.com/upload"
)

// Client is a Chrome Web Store API client.
type Client struct {
	// httpClient is the HTTP client used for API requests.
	httpClient *http.Client
	// baseURL is the base URL for API requests.
	baseURL string
	// uploadBaseURL is the base URL for upload requests.
	uploadBaseURL string

	// Publishers provides access to publishers resources.
	Publishers *PublishersService
	// Media provides access to media upload operations.
	Media *MediaService
}

// NewClient creates a new Chrome Web Store API client.
// The provided http.Client should be configured with OAuth 2.0 credentials.
func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	c := &Client{
		httpClient:    httpClient,
		baseURL:       DefaultBaseURL,
		uploadBaseURL: DefaultUploadBaseURL,
	}

	c.Publishers = newPublishersService(c)
	c.Media = newMediaService(c)

	return c
}

// SetBaseURL sets the base URL for API requests.
// This is useful for testing with a mock server.
func (c *Client) SetBaseURL(baseURL string) {
	c.baseURL = baseURL
}

// SetUploadBaseURL sets the base URL for upload requests.
// This is useful for testing with a mock server.
func (c *Client) SetUploadBaseURL(uploadBaseURL string) {
	c.uploadBaseURL = uploadBaseURL
}

// doRequest performs an HTTP request and returns the response.
func (c *Client) doRequest(ctx context.Context, method, urlStr string, body interface{}) (*http.Response, error) {
	var bodyReader io.Reader
	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("chromewebstore: failed to marshal request body: %w", err)
		}
		bodyReader = bytes.NewReader(jsonBody)
	}

	req, err := http.NewRequestWithContext(ctx, method, urlStr, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("chromewebstore: failed to create request: %w", err)
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")

	return c.httpClient.Do(req)
}

// doRequestWithMedia performs an HTTP request with media upload.
func (c *Client) doRequestWithMedia(ctx context.Context, method, urlStr string, media io.Reader, mediaType string) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, method, urlStr, media)
	if err != nil {
		return nil, fmt.Errorf("chromewebstore: failed to create request: %w", err)
	}

	if mediaType != "" {
		req.Header.Set("Content-Type", mediaType)
	}
	req.Header.Set("Accept", "application/json")

	return c.httpClient.Do(req)
}

// parseResponse parses the HTTP response into the target struct.
func parseResponse(resp *http.Response, target interface{}) error {
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("chromewebstore: failed to read response body: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		apiErr := newAPIError(resp, body)
		return apiErr
	}

	if target != nil && len(body) > 0 {
		if err := json.Unmarshal(body, target); err != nil {
			return fmt.Errorf("chromewebstore: failed to unmarshal response: %w", err)
		}
	}

	return nil
}

// buildURL builds a URL with the given base URL, path, and query parameters.
func buildURL(baseURL, path string, params url.Values) string {
	u := baseURL + path
	if len(params) > 0 {
		u += "?" + params.Encode()
	}
	return u
}
