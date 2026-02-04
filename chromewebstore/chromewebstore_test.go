package chromewebstore

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewClient(t *testing.T) {
	client := NewClient(nil)

	if client == nil {
		t.Fatal("expected client to be non-nil")
	}

	if client.Publishers == nil {
		t.Error("expected Publishers to be non-nil")
	}

	if client.Publishers.Items == nil {
		t.Error("expected Publishers.Items to be non-nil")
	}

	if client.Media == nil {
		t.Error("expected Media to be non-nil")
	}

	if client.baseURL != DefaultBaseURL {
		t.Errorf("expected baseURL to be %s, got %s", DefaultBaseURL, client.baseURL)
	}

	if client.uploadBaseURL != DefaultUploadBaseURL {
		t.Errorf("expected uploadBaseURL to be %s, got %s", DefaultUploadBaseURL, client.uploadBaseURL)
	}
}

func TestNewClientWithHTTPClient(t *testing.T) {
	httpClient := &http.Client{}
	client := NewClient(httpClient)

	if client == nil {
		t.Fatal("expected client to be non-nil")
	}

	if client.httpClient != httpClient {
		t.Error("expected httpClient to be the provided client")
	}
}

func TestSetBaseURL(t *testing.T) {
	client := NewClient(nil)
	newURL := "https://example.com"

	client.SetBaseURL(newURL)

	if client.baseURL != newURL {
		t.Errorf("expected baseURL to be %s, got %s", newURL, client.baseURL)
	}
}

func TestSetUploadBaseURL(t *testing.T) {
	client := NewClient(nil)
	newURL := "https://upload.example.com"

	client.SetUploadBaseURL(newURL)

	if client.uploadBaseURL != newURL {
		t.Errorf("expected uploadBaseURL to be %s, got %s", newURL, client.uploadBaseURL)
	}
}

func newTestServer(handler http.HandlerFunc) *httptest.Server {
	return httptest.NewServer(handler)
}
