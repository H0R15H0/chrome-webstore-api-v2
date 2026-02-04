// Package chromewebstore provides a client for the Chrome Web Store API v2.
package chromewebstore

import "fmt"

// ItemName represents a Chrome Web Store item resource name.
// Format: publishers/{publisherId}/items/{itemId}
type ItemName string

// NewItemName creates a new ItemName from publisher ID and item ID.
func NewItemName(publisherID, itemID string) ItemName {
	return ItemName(fmt.Sprintf("publishers/%s/items/%s", publisherID, itemID))
}

// String returns the string representation of ItemName.
func (n ItemName) String() string {
	return string(n)
}

// PublishTarget specifies the target for publishing.
type PublishTarget string

const (
	// PublishTargetDefault publishes to the default audience.
	PublishTargetDefault PublishTarget = "PUBLISH_TARGET_UNSPECIFIED"
	// PublishTargetTrustedTesters publishes to trusted testers only.
	PublishTargetTrustedTesters PublishTarget = "TRUSTED_TESTERS"
)

// ItemState represents the state of a Chrome Web Store item.
type ItemState string

const (
	ItemStateUnspecified    ItemState = "STATE_UNSPECIFIED"
	ItemStateDraft          ItemState = "DRAFT"
	ItemStatePendingReview  ItemState = "PENDING_REVIEW"
	ItemStatePublished      ItemState = "PUBLISHED"
	ItemStateRejected       ItemState = "REJECTED"
	ItemStateTakenDown      ItemState = "TAKEN_DOWN"
	ItemStateNotPublished   ItemState = "NOT_PUBLISHED"
	ItemStateSuspended      ItemState = "SUSPENDED"
	ItemStateInReview       ItemState = "IN_REVIEW"
	ItemStatePendingPublish ItemState = "PENDING_PUBLISH"
)

// StatusCode represents the status code in API responses.
type StatusCode string

const (
	StatusCodeUnspecified       StatusCode = "STATUS_CODE_UNSPECIFIED"
	StatusCodeOK                StatusCode = "OK"
	StatusCodeInvalidItem       StatusCode = "INVALID_ITEM"
	StatusCodeUnauthorized      StatusCode = "UNAUTHORIZED"
	StatusCodeNotFound          StatusCode = "NOT_FOUND"
	StatusCodeConflict          StatusCode = "CONFLICT"
	StatusCodeInternalError     StatusCode = "INTERNAL_ERROR"
	StatusCodeInvalidDeveloper  StatusCode = "INVALID_DEVELOPER"
	StatusCodeRateLimitExceeded StatusCode = "RATE_LIMIT_EXCEEDED"
)

// ItemStatus represents the response from fetchStatus API.
type ItemStatus struct {
	// Name is the resource name of the item.
	Name string `json:"name,omitempty"`
	// State is the current state of the item.
	State ItemState `json:"state,omitempty"`
	// Version is the current version of the item.
	Version string `json:"version,omitempty"`
	// DetailedStatus provides detailed status information.
	DetailedStatus []StatusDetail `json:"detailedStatus,omitempty"`
	// DownloadURL is the URL to download the item (if available).
	DownloadURL string `json:"downloadUrl,omitempty"`
}

// StatusDetail provides detailed status information.
type StatusDetail struct {
	// StatusCode is the status code.
	StatusCode StatusCode `json:"statusCode,omitempty"`
	// Description provides a human-readable description.
	Description string `json:"description,omitempty"`
}

// PublishRequest represents a request to publish an item.
type PublishRequest struct {
	// PublishTarget specifies the target audience.
	PublishTarget PublishTarget `json:"publishTarget,omitempty"`
}

// PublishResponse represents the response from publish API.
type PublishResponse struct {
	// StatusCode is the status of the publish operation.
	StatusCode StatusCode `json:"statusCode,omitempty"`
	// StatusDetail provides additional status details.
	StatusDetail []string `json:"statusDetail,omitempty"`
}

// UploadResponse represents the response from upload API.
type UploadResponse struct {
	// Name is the resource name of the uploaded item.
	Name string `json:"name,omitempty"`
	// StatusCode is the status of the upload operation.
	StatusCode StatusCode `json:"statusCode,omitempty"`
	// StatusDetail provides additional status details.
	StatusDetail []string `json:"statusDetail,omitempty"`
	// ItemError contains error details if the upload failed.
	ItemError []ItemError `json:"itemError,omitempty"`
}

// ItemError represents an error related to an item.
type ItemError struct {
	// ErrorCode is the error code.
	ErrorCode string `json:"errorCode,omitempty"`
	// ErrorDetail provides a human-readable error description.
	ErrorDetail string `json:"errorDetail,omitempty"`
}

// SetPublishedDeployPercentageRequest represents a request to set deploy percentage.
type SetPublishedDeployPercentageRequest struct {
	// DeployPercentage is the percentage of users to deploy to (0-100).
	DeployPercentage int `json:"deployPercentage"`
}

// SetPublishedDeployPercentageResponse represents the response from setPublishedDeployPercentage API.
type SetPublishedDeployPercentageResponse struct {
	// StatusCode is the status of the operation.
	StatusCode StatusCode `json:"statusCode,omitempty"`
	// StatusDetail provides additional status details.
	StatusDetail []string `json:"statusDetail,omitempty"`
}

// CancelSubmissionResponse represents the response from cancelSubmission API.
type CancelSubmissionResponse struct {
	// StatusCode is the status of the operation.
	StatusCode StatusCode `json:"statusCode,omitempty"`
	// StatusDetail provides additional status details.
	StatusDetail []string `json:"statusDetail,omitempty"`
}
