package chromewebstore

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// SetPublishedDeployPercentageCall represents a call to set the deploy percentage.
type SetPublishedDeployPercentageCall struct {
	client           *Client
	name             ItemName
	ctx              context.Context
	params           url.Values
	deployPercentage int
}

// newSetPublishedDeployPercentageCall creates a new SetPublishedDeployPercentageCall.
func newSetPublishedDeployPercentageCall(c *Client, name ItemName) *SetPublishedDeployPercentageCall {
	return &SetPublishedDeployPercentageCall{
		client:           c,
		name:             name,
		ctx:              context.Background(),
		params:           make(url.Values),
		deployPercentage: 100,
	}
}

// Context sets the context for the request.
func (c *SetPublishedDeployPercentageCall) Context(ctx context.Context) *SetPublishedDeployPercentageCall {
	c.ctx = ctx
	return c
}

// DeployPercentage sets the deploy percentage (0-100).
func (c *SetPublishedDeployPercentageCall) DeployPercentage(percent int) *SetPublishedDeployPercentageCall {
	c.deployPercentage = percent
	return c
}

// Do executes the set deploy percentage request.
func (c *SetPublishedDeployPercentageCall) Do() (*SetPublishedDeployPercentageResponse, error) {
	path := fmt.Sprintf("/v2/%s:setPublishedDeployPercentage", c.name)

	c.params.Set("deployPercentage", fmt.Sprintf("%d", c.deployPercentage))

	urlStr := buildURL(c.client.baseURL, path, c.params)

	resp, err := c.client.doRequest(c.ctx, http.MethodPost, urlStr, nil)
	if err != nil {
		return nil, err
	}

	var result SetPublishedDeployPercentageResponse
	if err := parseResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
