package chromewebstore

import (
	"context"
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

const (
	// ScopeChromeWebStore is the OAuth 2.0 scope for Chrome Web Store API.
	ScopeChromeWebStore = "https://www.googleapis.com/auth/chromewebstore"
	// ScopeChromeWebStoreReadOnly is the read-only OAuth 2.0 scope.
	ScopeChromeWebStoreReadOnly = "https://www.googleapis.com/auth/chromewebstore.readonly"
)

// AuthConfig holds the OAuth 2.0 configuration for authentication.
type AuthConfig struct {
	// ClientID is the OAuth 2.0 client ID.
	ClientID string
	// ClientSecret is the OAuth 2.0 client secret.
	ClientSecret string
	// RefreshToken is the OAuth 2.0 refresh token.
	RefreshToken string
}

// NewAuthenticatedClient creates a new HTTP client with OAuth 2.0 authentication.
// The client will automatically refresh the access token when needed.
func NewAuthenticatedClient(ctx context.Context, config AuthConfig) *http.Client {
	oauthConfig := &oauth2.Config{
		ClientID:     config.ClientID,
		ClientSecret: config.ClientSecret,
		Endpoint:     google.Endpoint,
		Scopes:       []string{ScopeChromeWebStore},
	}

	token := &oauth2.Token{
		RefreshToken: config.RefreshToken,
	}

	return oauthConfig.Client(ctx, token)
}

// NewClientFromCredentials creates a new Chrome Web Store API client
// with OAuth 2.0 authentication using the provided credentials.
func NewClientFromCredentials(ctx context.Context, config AuthConfig) *Client {
	httpClient := NewAuthenticatedClient(ctx, config)
	return NewClient(httpClient)
}

// NewClientFromAccessToken creates a new Chrome Web Store API client
// using a static access token.
// Note: Access tokens expire after 1 hour and cannot be refreshed with this method.
func NewClientFromAccessToken(accessToken string) *Client {
	token := &oauth2.Token{
		AccessToken: accessToken,
	}
	httpClient := oauth2.NewClient(context.Background(), oauth2.StaticTokenSource(token))
	return NewClient(httpClient)
}
