package chromewebstore

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// PublishCall represents a call to publish an item.
type PublishCall struct {
	client  *Client
	name    ItemName
	ctx     context.Context
	params  url.Values
	request *PublishRequest
}

// newPublishCall creates a new PublishCall.
func newPublishCall(c *Client, name ItemName) *PublishCall {
	return &PublishCall{
		client:  c,
		name:    name,
		ctx:     context.Background(),
		params:  make(url.Values),
		request: &PublishRequest{},
	}
}

// Context sets the context for the request.
func (c *PublishCall) Context(ctx context.Context) *PublishCall {
	c.ctx = ctx
	return c
}

// PublishType sets the publish type (IMMEDIATE or STAGED).
func (c *PublishCall) PublishType(publishType PublishType) *PublishCall {
	c.request.PublishType = publishType
	return c
}

// SkipReview sets whether to attempt bypassing review.
func (c *PublishCall) SkipReview(skip bool) *PublishCall {
	c.request.SkipReview = skip
	return c
}

// DeployPercentage sets the deploy percentage for staged rollout.
func (c *PublishCall) DeployPercentage(percentage int) *PublishCall {
	c.request.DeployInfos = []DeployInfo{{DeployPercentage: percentage}}
	return c
}

// Do executes the publish request.
func (c *PublishCall) Do() (*PublishResponse, error) {
	path := fmt.Sprintf("/v2/%s:publish", c.name)
	urlStr := buildURL(c.client.baseURL, path, c.params)

	resp, err := c.client.doRequest(c.ctx, http.MethodPost, urlStr, c.request)
	if err != nil {
		return nil, err
	}

	var result PublishResponse
	if err := parseResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
