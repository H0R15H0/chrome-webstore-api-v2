package chromewebstore

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// FetchStatusCall represents a call to fetch the status of an item.
type FetchStatusCall struct {
	client  *Client
	name    ItemName
	ctx     context.Context
	params  url.Values
}

// newFetchStatusCall creates a new FetchStatusCall.
func newFetchStatusCall(c *Client, name ItemName) *FetchStatusCall {
	return &FetchStatusCall{
		client: c,
		name:   name,
		ctx:    context.Background(),
		params: make(url.Values),
	}
}

// Context sets the context for the request.
func (c *FetchStatusCall) Context(ctx context.Context) *FetchStatusCall {
	c.ctx = ctx
	return c
}

// Projection sets the projection parameter.
// Valid values are "DRAFT" or "PUBLISHED".
func (c *FetchStatusCall) Projection(projection string) *FetchStatusCall {
	c.params.Set("projection", projection)
	return c
}

// Do executes the fetch status request.
func (c *FetchStatusCall) Do() (*ItemStatus, error) {
	path := fmt.Sprintf("/v2/%s:fetchStatus", c.name)
	urlStr := buildURL(c.client.baseURL, path, c.params)

	resp, err := c.client.doRequest(c.ctx, http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, err
	}

	var result ItemStatus
	if err := parseResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
