package chromewebstore

import (
	"context"
	"testing"
)

func TestNewAuthenticatedClient(t *testing.T) {
	config := AuthConfig{
		ClientID:     "test-client-id",
		ClientSecret: "test-client-secret",
		RefreshToken: "test-refresh-token",
	}

	client := NewAuthenticatedClient(context.Background(), config)

	if client == nil {
		t.Fatal("expected client to be non-nil")
	}
}

func TestNewClientFromCredentials(t *testing.T) {
	config := AuthConfig{
		ClientID:     "test-client-id",
		ClientSecret: "test-client-secret",
		RefreshToken: "test-refresh-token",
	}

	client := NewClientFromCredentials(context.Background(), config)

	if client == nil {
		t.Fatal("expected client to be non-nil")
	}

	if client.Publishers == nil {
		t.Error("expected Publishers to be non-nil")
	}

	if client.Media == nil {
		t.Error("expected Media to be non-nil")
	}
}
