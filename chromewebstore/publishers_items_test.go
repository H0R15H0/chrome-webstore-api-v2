package chromewebstore

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"testing"
)

func TestFetchStatus(t *testing.T) {
	expectedResponse := ItemStatus{
		Name:   "publishers/test-publisher/items/test-item",
		ItemID: "test-item",
		PublishedItemRevisionStatus: &ItemRevisionStatus{
			State: ItemStatePublished,
			DistributionChannels: []DistributionChannel{
				{DeployPercentage: 100, CrxVersion: "1.0.0"},
			},
		},
	}

	server := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("expected GET method, got %s", r.Method)
		}

		if !strings.Contains(r.URL.Path, ":fetchStatus") {
			t.Errorf("expected path to contain :fetchStatus, got %s", r.URL.Path)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(expectedResponse)
	})
	defer server.Close()

	client := NewClient(nil)
	client.SetBaseURL(server.URL)

	itemName := NewItemName("test-publisher", "test-item")
	status, err := client.Publishers.Items.FetchStatus(itemName).Context(context.Background()).Do()

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if status.Name != expectedResponse.Name {
		t.Errorf("expected name %s, got %s", expectedResponse.Name, status.Name)
	}

	if status.PublishedItemRevisionStatus == nil {
		t.Fatal("expected PublishedItemRevisionStatus to be non-nil")
	}

	if status.PublishedItemRevisionStatus.State != expectedResponse.PublishedItemRevisionStatus.State {
		t.Errorf("expected state %s, got %s", expectedResponse.PublishedItemRevisionStatus.State, status.PublishedItemRevisionStatus.State)
	}

	if len(status.PublishedItemRevisionStatus.DistributionChannels) == 0 {
		t.Fatal("expected at least one distribution channel")
	}

	if status.PublishedItemRevisionStatus.DistributionChannels[0].CrxVersion != "1.0.0" {
		t.Errorf("expected version 1.0.0, got %s", status.PublishedItemRevisionStatus.DistributionChannels[0].CrxVersion)
	}
}

func TestFetchStatusWithProjection(t *testing.T) {
	server := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		projection := r.URL.Query().Get("projection")
		if projection != "DRAFT" {
			t.Errorf("expected projection DRAFT, got %s", projection)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(ItemStatus{})
	})
	defer server.Close()

	client := NewClient(nil)
	client.SetBaseURL(server.URL)

	itemName := NewItemName("test-publisher", "test-item")
	_, err := client.Publishers.Items.FetchStatus(itemName).Projection("DRAFT").Do()

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestPublish(t *testing.T) {
	expectedResponse := PublishResponse{
		Name:   "publishers/test-publisher/items/test-item",
		ItemID: "test-item",
		State:  ItemStatePendingReview,
	}

	server := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST method, got %s", r.Method)
		}

		if !strings.Contains(r.URL.Path, ":publish") {
			t.Errorf("expected path to contain :publish, got %s", r.URL.Path)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(expectedResponse)
	})
	defer server.Close()

	client := NewClient(nil)
	client.SetBaseURL(server.URL)

	itemName := NewItemName("test-publisher", "test-item")
	resp, err := client.Publishers.Items.Publish(itemName).Context(context.Background()).Do()

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if resp.State != expectedResponse.State {
		t.Errorf("expected state %s, got %s", expectedResponse.State, resp.State)
	}
}

func TestPublishWithType(t *testing.T) {
	server := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		var req PublishRequest
		json.NewDecoder(r.Body).Decode(&req)
		if req.PublishType != PublishTypeImmediate {
			t.Errorf("expected publishType %s, got %s", PublishTypeImmediate, req.PublishType)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(PublishResponse{})
	})
	defer server.Close()

	client := NewClient(nil)
	client.SetBaseURL(server.URL)

	itemName := NewItemName("test-publisher", "test-item")
	_, err := client.Publishers.Items.Publish(itemName).PublishType(PublishTypeImmediate).Do()

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestCancelSubmission(t *testing.T) {
	server := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST method, got %s", r.Method)
		}

		if !strings.Contains(r.URL.Path, ":cancelSubmission") {
			t.Errorf("expected path to contain :cancelSubmission, got %s", r.URL.Path)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(CancelSubmissionResponse{})
	})
	defer server.Close()

	client := NewClient(nil)
	client.SetBaseURL(server.URL)

	itemName := NewItemName("test-publisher", "test-item")
	_, err := client.Publishers.Items.CancelSubmission(itemName).Context(context.Background()).Do()

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestSetPublishedDeployPercentage(t *testing.T) {
	server := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST method, got %s", r.Method)
		}

		if !strings.Contains(r.URL.Path, ":setPublishedDeployPercentage") {
			t.Errorf("expected path to contain :setPublishedDeployPercentage, got %s", r.URL.Path)
		}

		percentage := r.URL.Query().Get("deployPercentage")
		if percentage != "50" {
			t.Errorf("expected deployPercentage 50, got %s", percentage)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(SetPublishedDeployPercentageResponse{})
	})
	defer server.Close()

	client := NewClient(nil)
	client.SetBaseURL(server.URL)

	itemName := NewItemName("test-publisher", "test-item")
	_, err := client.Publishers.Items.SetPublishedDeployPercentage(itemName).DeployPercentage(50).Do()

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestAPIError(t *testing.T) {
	server := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"error": "not found"}`))
	})
	defer server.Close()

	client := NewClient(nil)
	client.SetBaseURL(server.URL)

	itemName := NewItemName("test-publisher", "test-item")
	_, err := client.Publishers.Items.FetchStatus(itemName).Do()

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	apiErr, ok := err.(*APIError)
	if !ok {
		t.Fatalf("expected *APIError, got %T", err)
	}

	if !apiErr.IsNotFound() {
		t.Error("expected IsNotFound to be true")
	}
}

func TestItemName(t *testing.T) {
	name := NewItemName("my-publisher", "my-item")
	expected := "publishers/my-publisher/items/my-item"

	if name.String() != expected {
		t.Errorf("expected %s, got %s", expected, name.String())
	}
}
