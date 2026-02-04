package chromewebstore

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// PublishCall represents a call to publish an item.
type PublishCall struct {
	client        *Client
	name          ItemName
	ctx           context.Context
	params        url.Values
	publishTarget PublishTarget
}

// newPublishCall creates a new PublishCall.
func newPublishCall(c *Client, name ItemName) *PublishCall {
	return &PublishCall{
		client: c,
		name:   name,
		ctx:    context.Background(),
		params: make(url.Values),
	}
}

// Context sets the context for the request.
func (c *PublishCall) Context(ctx context.Context) *PublishCall {
	c.ctx = ctx
	return c
}

// PublishTarget sets the publish target.
// Use PublishTargetTrustedTesters to publish to trusted testers only.
func (c *PublishCall) PublishTarget(target PublishTarget) *PublishCall {
	c.publishTarget = target
	return c
}

// Do executes the publish request.
func (c *PublishCall) Do() (*PublishResponse, error) {
	path := fmt.Sprintf("/v2/%s:publish", c.name)

	if c.publishTarget != "" {
		c.params.Set("publishTarget", string(c.publishTarget))
	}

	urlStr := buildURL(c.client.baseURL, path, c.params)

	resp, err := c.client.doRequest(c.ctx, http.MethodPost, urlStr, nil)
	if err != nil {
		return nil, err
	}

	var result PublishResponse
	if err := parseResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
