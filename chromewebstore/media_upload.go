package chromewebstore

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// UploadCall represents a call to upload an extension package.
type UploadCall struct {
	client    *Client
	name      ItemName
	ctx       context.Context
	params    url.Values
	media     io.Reader
	mediaType string
}

// newUploadCall creates a new UploadCall.
func newUploadCall(c *Client, name ItemName) *UploadCall {
	return &UploadCall{
		client:    c,
		name:      name,
		ctx:       context.Background(),
		params:    make(url.Values),
		mediaType: "application/zip",
	}
}

// Context sets the context for the request.
func (c *UploadCall) Context(ctx context.Context) *UploadCall {
	c.ctx = ctx
	return c
}

// Media sets the media to upload and its content type.
func (c *UploadCall) Media(media io.Reader, mediaType string) *UploadCall {
	c.media = media
	if mediaType != "" {
		c.mediaType = mediaType
	}
	return c
}

// Do executes the upload request.
func (c *UploadCall) Do() (*UploadResponse, error) {
	if c.media == nil {
		return nil, fmt.Errorf("chromewebstore: media is required for upload")
	}

	path := fmt.Sprintf("/v2/%s:upload", c.name)
	urlStr := buildURL(c.client.uploadBaseURL, path, c.params)

	resp, err := c.client.doRequestWithMedia(c.ctx, http.MethodPost, urlStr, c.media, c.mediaType)
	if err != nil {
		return nil, err
	}

	var result UploadResponse
	if err := parseResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
