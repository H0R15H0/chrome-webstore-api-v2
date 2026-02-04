package chromewebstore

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// CancelSubmissionCall represents a call to cancel a pending submission.
type CancelSubmissionCall struct {
	client *Client
	name   ItemName
	ctx    context.Context
	params url.Values
}

// newCancelSubmissionCall creates a new CancelSubmissionCall.
func newCancelSubmissionCall(c *Client, name ItemName) *CancelSubmissionCall {
	return &CancelSubmissionCall{
		client: c,
		name:   name,
		ctx:    context.Background(),
		params: make(url.Values),
	}
}

// Context sets the context for the request.
func (c *CancelSubmissionCall) Context(ctx context.Context) *CancelSubmissionCall {
	c.ctx = ctx
	return c
}

// Do executes the cancel submission request.
func (c *CancelSubmissionCall) Do() (*CancelSubmissionResponse, error) {
	path := fmt.Sprintf("/v2/%s:cancelSubmission", c.name)
	urlStr := buildURL(c.client.baseURL, path, c.params)

	resp, err := c.client.doRequest(c.ctx, http.MethodPost, urlStr, nil)
	if err != nil {
		return nil, err
	}

	var result CancelSubmissionResponse
	if err := parseResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
