// This file was auto-generated by Fern from our API Definition.

package api

import (
	core "github.com/hookdeck/hookdeck-go-sdk/core"
	time "time"
)

type EventBulkRetryCreateRequest struct {
	// Filter properties for the events to be included in the bulk retry
	Query *core.Optional[EventBulkRetryCreateRequestQuery] `json:"query,omitempty"`
}

type EventBulkRetryListRequest struct {
	CancelledAt       *time.Time                    `json:"-"`
	CompletedAt       *time.Time                    `json:"-"`
	CreatedAt         *time.Time                    `json:"-"`
	Id                *string                       `json:"-"`
	QueryPartialMatch *bool                         `json:"-"`
	InProgress        *bool                         `json:"-"`
	OrderBy           *string                       `json:"-"`
	Dir               *EventBulkRetryListRequestDir `json:"-"`
	Limit             *int                          `json:"-"`
	Next              *string                       `json:"-"`
	Prev              *string                       `json:"-"`
}
