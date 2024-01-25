// This file was auto-generated by Fern from our API Definition.

package api

import (
	fmt "fmt"
	core "github.com/hookdeck/hookdeck-go-sdk/core"
	time "time"
)

type IgnoredEventBulkRetryCreateRequest struct {
	// Filter by the bulk retry ignored event query object
	Query *core.Optional[IgnoredEventBulkRetryCreateRequestQuery] `json:"query,omitempty"`
}

type IgnoredEventBulkRetryListRequest struct {
	CancelledAt       *time.Time                               `json:"-"`
	CompletedAt       *time.Time                               `json:"-"`
	CreatedAt         *time.Time                               `json:"-"`
	Id                *string                                  `json:"-"`
	QueryPartialMatch *bool                                    `json:"-"`
	InProgress        *bool                                    `json:"-"`
	OrderBy           *IgnoredEventBulkRetryListRequestOrderBy `json:"-"`
	Dir               *IgnoredEventBulkRetryListRequestDir     `json:"-"`
	Limit             *int                                     `json:"-"`
	Next              *string                                  `json:"-"`
	Prev              *string                                  `json:"-"`
}

// Filter by the bulk retry ignored event query object
type IgnoredEventBulkRetryCreateRequestQuery struct {
	// The cause of the ignored event
	Cause *IgnoredEventBulkRetryCreateRequestQueryCause `json:"cause,omitempty"`
	// Connection ID of the ignored event
	WebhookId *IgnoredEventBulkRetryCreateRequestQueryWebhookId `json:"webhook_id,omitempty"`
	// The associated transformation ID (only applicable to the cause `TRANSFORMATION_FAILED`)
	TransformationId *string `json:"transformation_id,omitempty"`
}

type IgnoredEventBulkRetryListRequestDir string

const (
	IgnoredEventBulkRetryListRequestDirAsc  IgnoredEventBulkRetryListRequestDir = "asc"
	IgnoredEventBulkRetryListRequestDirDesc IgnoredEventBulkRetryListRequestDir = "desc"
)

func NewIgnoredEventBulkRetryListRequestDirFromString(s string) (IgnoredEventBulkRetryListRequestDir, error) {
	switch s {
	case "asc":
		return IgnoredEventBulkRetryListRequestDirAsc, nil
	case "desc":
		return IgnoredEventBulkRetryListRequestDirDesc, nil
	}
	var t IgnoredEventBulkRetryListRequestDir
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (i IgnoredEventBulkRetryListRequestDir) Ptr() *IgnoredEventBulkRetryListRequestDir {
	return &i
}

type IgnoredEventBulkRetryListRequestOrderBy string

const (
	IgnoredEventBulkRetryListRequestOrderByCreatedAt IgnoredEventBulkRetryListRequestOrderBy = "created_at"
)

func NewIgnoredEventBulkRetryListRequestOrderByFromString(s string) (IgnoredEventBulkRetryListRequestOrderBy, error) {
	switch s {
	case "created_at":
		return IgnoredEventBulkRetryListRequestOrderByCreatedAt, nil
	}
	var t IgnoredEventBulkRetryListRequestOrderBy
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (i IgnoredEventBulkRetryListRequestOrderBy) Ptr() *IgnoredEventBulkRetryListRequestOrderBy {
	return &i
}

type IgnoredEventBulkRetryPlanResponse struct {
	// Number of batches required to complete the bulk retry
	EstimatedBatch *int `json:"estimated_batch,omitempty"`
	// Number of estimated events to be retried
	EstimatedCount *int `json:"estimated_count,omitempty"`
	// Progression of the batch operations, values 0 - 1
	Progress *float64 `json:"progress,omitempty"`
}