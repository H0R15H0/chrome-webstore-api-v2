package chromewebstore

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"testing"
)

func TestUpload(t *testing.T) {
	expectedResponse := UploadResponse{
		Name:        "publishers/test-publisher/items/test-item",
		ItemID:      "test-item",
		UploadState: UploadStateSucceeded,
		CrxVersion:  "1.0.0",
	}

	server := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST method, got %s", r.Method)
		}

		if !strings.Contains(r.URL.Path, ":upload") {
			t.Errorf("expected path to contain :upload, got %s", r.URL.Path)
		}

		contentType := r.Header.Get("Content-Type")
		if contentType != "application/zip" {
			t.Errorf("expected Content-Type application/zip, got %s", contentType)
		}

		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatalf("failed to read body: %v", err)
		}
		if string(body) != "test-content" {
			t.Errorf("expected body 'test-content', got %s", string(body))
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(expectedResponse)
	})
	defer server.Close()

	client := NewClient(nil)
	client.SetUploadBaseURL(server.URL)

	itemName := NewItemName("test-publisher", "test-item")
	media := strings.NewReader("test-content")

	resp, err := client.Media.Upload(itemName).Context(context.Background()).Media(media, "application/zip").Do()

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if resp.Name != expectedResponse.Name {
		t.Errorf("expected name %s, got %s", expectedResponse.Name, resp.Name)
	}

	if resp.UploadState != expectedResponse.UploadState {
		t.Errorf("expected upload state %s, got %s", expectedResponse.UploadState, resp.UploadState)
	}
}

func TestUploadWithoutMedia(t *testing.T) {
	client := NewClient(nil)

	itemName := NewItemName("test-publisher", "test-item")
	_, err := client.Media.Upload(itemName).Do()

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if !strings.Contains(err.Error(), "media is required") {
		t.Errorf("expected error about media being required, got %v", err)
	}
}

func TestUploadWithCustomMediaType(t *testing.T) {
	server := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		contentType := r.Header.Get("Content-Type")
		if contentType != "application/octet-stream" {
			t.Errorf("expected Content-Type application/octet-stream, got %s", contentType)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(UploadResponse{})
	})
	defer server.Close()

	client := NewClient(nil)
	client.SetUploadBaseURL(server.URL)

	itemName := NewItemName("test-publisher", "test-item")
	media := strings.NewReader("test-content")

	_, err := client.Media.Upload(itemName).Media(media, "application/octet-stream").Do()

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}
