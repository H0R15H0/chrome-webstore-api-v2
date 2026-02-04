package chromewebstore

// MediaService provides access to media upload operations.
type MediaService struct {
	client *Client
}

// newMediaService creates a new MediaService.
func newMediaService(c *Client) *MediaService {
	return &MediaService{client: c}
}

// Upload returns an UploadCall for uploading an extension package.
func (s *MediaService) Upload(name ItemName) *UploadCall {
	return newUploadCall(s.client, name)
}
