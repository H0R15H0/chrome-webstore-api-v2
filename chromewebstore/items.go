package chromewebstore

// PublishersService provides access to publishers resources.
type PublishersService struct {
	client *Client
	// Items provides access to items operations.
	Items *ItemsService
}

// newPublishersService creates a new PublishersService.
func newPublishersService(c *Client) *PublishersService {
	s := &PublishersService{client: c}
	s.Items = newItemsService(c)
	return s
}

// ItemsService provides access to items operations.
type ItemsService struct {
	client *Client
}

// newItemsService creates a new ItemsService.
func newItemsService(c *Client) *ItemsService {
	return &ItemsService{client: c}
}

// FetchStatus returns a FetchStatusCall for fetching the status of an item.
func (s *ItemsService) FetchStatus(name ItemName) *FetchStatusCall {
	return newFetchStatusCall(s.client, name)
}

// Publish returns a PublishCall for publishing an item.
func (s *ItemsService) Publish(name ItemName) *PublishCall {
	return newPublishCall(s.client, name)
}

// CancelSubmission returns a CancelSubmissionCall for canceling a pending submission.
func (s *ItemsService) CancelSubmission(name ItemName) *CancelSubmissionCall {
	return newCancelSubmissionCall(s.client, name)
}

// SetPublishedDeployPercentage returns a SetPublishedDeployPercentageCall
// for setting the deploy percentage of a published item.
func (s *ItemsService) SetPublishedDeployPercentage(name ItemName) *SetPublishedDeployPercentageCall {
	return newSetPublishedDeployPercentageCall(s.client, name)
}
