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

// PublishType specifies the type of publishing.
type PublishType string

const (
	// PublishTypeUnspecified is the default publish type.
	PublishTypeUnspecified PublishType = "PUBLISH_TYPE_UNSPECIFIED"
	// PublishTypeDefault publishes immediately after approval.
	PublishTypeDefault PublishType = "DEFAULT_PUBLISH"
	// PublishTypeStaged stages after approval, allowing manual publish later.
	PublishTypeStaged PublishType = "STAGED_PUBLISH"
)

// ItemState represents the state of a Chrome Web Store item revision.
type ItemState string

const (
	ItemStateUnspecified       ItemState = "STATE_UNSPECIFIED"
	ItemStatePendingReview     ItemState = "PENDING_REVIEW"
	ItemStateStaged            ItemState = "STAGED"
	ItemStatePublished         ItemState = "PUBLISHED"
	ItemStatePublishedToTesters ItemState = "PUBLISHED_TO_TESTERS"
	ItemStateRejected          ItemState = "REJECTED"
	ItemStateCancelled         ItemState = "CANCELLED"
)

// UploadState represents the state of an upload operation.
type UploadState string

const (
	UploadStateUnspecified UploadState = "UPLOAD_STATE_UNSPECIFIED"
	UploadStateSucceeded   UploadState = "SUCCEEDED"
	UploadStateInProgress  UploadState = "IN_PROGRESS"
	UploadStateFailed      UploadState = "FAILED"
	UploadStateNotFound    UploadState = "NOT_FOUND"
)

// ItemStatus represents the response from fetchStatus API (FetchItemStatusResponse).
type ItemStatus struct {
	// Name is the resource name of the item.
	Name string `json:"name,omitempty"`
	// ItemID is the ID of the item.
	ItemID string `json:"itemId,omitempty"`
	// PublicKey is the public key of the item.
	PublicKey string `json:"publicKey,omitempty"`
	// Warned indicates if there is a policy violation warning.
	Warned bool `json:"warned,omitempty"`
	// TakenDown indicates if the item was taken down due to policy violation.
	TakenDown bool `json:"takenDown,omitempty"`
	// LastAsyncUploadState is the state of the most recent async upload.
	LastAsyncUploadState UploadState `json:"lastAsyncUploadState,omitempty"`
	// SubmittedItemRevisionStatus contains the status of a pending submission.
	SubmittedItemRevisionStatus *ItemRevisionStatus `json:"submittedItemRevisionStatus,omitempty"`
	// PublishedItemRevisionStatus contains the status of the published revision.
	PublishedItemRevisionStatus *ItemRevisionStatus `json:"publishedItemRevisionStatus,omitempty"`
}

// ItemRevisionStatus represents the status of an item revision.
type ItemRevisionStatus struct {
	// State is the current state of the item revision.
	State ItemState `json:"state,omitempty"`
	// DistributionChannels contains distribution information.
	DistributionChannels []DistributionChannel `json:"distributionChannels,omitempty"`
}

// DistributionChannel represents a distribution channel for an item.
type DistributionChannel struct {
	// DeployPercentage is the percentage of users receiving this version (0-100).
	DeployPercentage int `json:"deployPercentage,omitempty"`
	// CrxVersion is the version string of the CRX from the manifest.
	CrxVersion string `json:"crxVersion,omitempty"`
}

// PublishRequest represents a request to publish an item (PublishItemRequest).
type PublishRequest struct {
	// PublishType controls immediate vs staged publishing.
	PublishType PublishType `json:"publishType,omitempty"`
	// SkipReview indicates whether to attempt bypassing review.
	SkipReview bool `json:"skipReview,omitempty"`
	// DeployInfos contains optional deployment parameters.
	DeployInfos []DeployInfo `json:"deployInfos,omitempty"`
}

// DeployInfo represents deployment configuration for a release channel.
type DeployInfo struct {
	// DeployPercentage is the percentage for rollout (0-100).
	DeployPercentage int `json:"deployPercentage,omitempty"`
}

// PublishResponse represents the response from publish API (PublishItemResponse).
type PublishResponse struct {
	// Name is the resource name of the item.
	Name string `json:"name,omitempty"`
	// ItemID is the ID of the item.
	ItemID string `json:"itemId,omitempty"`
	// State is the current submission status.
	State ItemState `json:"state,omitempty"`
}

// UploadResponse represents the response from upload API (UploadItemPackageResponse).
type UploadResponse struct {
	// Name is the resource name of the target item.
	Name string `json:"name,omitempty"`
	// ItemID is the ID of the item that received the package.
	ItemID string `json:"itemId,omitempty"`
	// UploadState is the status of the upload.
	UploadState UploadState `json:"uploadState,omitempty"`
	// CrxVersion is the extension version from the manifest (unset if upload in progress).
	CrxVersion string `json:"crxVersion,omitempty"`
}

// SetPublishedDeployPercentageRequest represents a request to set deploy percentage.
type SetPublishedDeployPercentageRequest struct {
	// DeployPercentage is the percentage of users to deploy to (0-100).
	// Must exceed the current value.
	DeployPercentage int `json:"deployPercentage"`
}

// SetPublishedDeployPercentageResponse represents the response from setPublishedDeployPercentage API.
// This is an empty response according to API spec.
type SetPublishedDeployPercentageResponse struct{}

// CancelSubmissionResponse represents the response from cancelSubmission API.
// This is an empty response according to API spec.
type CancelSubmissionResponse struct{}
