// This file was auto-generated by Fern from our API Definition.

package api

import (
	fmt "fmt"
	time "time"
)

type RequestListRequest struct {
	Id             []*string                  `json:"-" url:"id,omitempty"`
	Status         *RequestListRequestStatus  `json:"-" url:"status,omitempty"`
	RejectionCause *RequestRejectionCause     `json:"-" url:"rejection_cause,omitempty"`
	SourceId       []*string                  `json:"-" url:"source_id,omitempty"`
	Verified       *bool                      `json:"-" url:"verified,omitempty"`
	SearchTerm     *string                    `json:"-" url:"search_term,omitempty"`
	Headers        *string                    `json:"-" url:"headers,omitempty"`
	Body           *string                    `json:"-" url:"body,omitempty"`
	ParsedQuery    *string                    `json:"-" url:"parsed_query,omitempty"`
	Path           *string                    `json:"-" url:"path,omitempty"`
	IgnoredCount   *int                       `json:"-" url:"ignored_count,omitempty"`
	EventsCount    *int                       `json:"-" url:"events_count,omitempty"`
	IngestedAt     *time.Time                 `json:"-" url:"ingested_at,omitempty"`
	BulkRetryId    []*string                  `json:"-" url:"bulk_retry_id,omitempty"`
	Include        *string                    `json:"-" url:"include,omitempty"`
	OrderBy        *RequestListRequestOrderBy `json:"-" url:"order_by,omitempty"`
	Dir            *RequestListRequestDir     `json:"-" url:"dir,omitempty"`
	Limit          *int                       `json:"-" url:"limit,omitempty"`
	Next           *string                    `json:"-" url:"next,omitempty"`
	Prev           *string                    `json:"-" url:"prev,omitempty"`
}

type RequestListEventRequest struct {
	Id             []*string                       `json:"-" url:"id,omitempty"`
	Status         *EventStatus                    `json:"-" url:"status,omitempty"`
	WebhookId      []*string                       `json:"-" url:"webhook_id,omitempty"`
	DestinationId  []*string                       `json:"-" url:"destination_id,omitempty"`
	SourceId       []*string                       `json:"-" url:"source_id,omitempty"`
	Attempts       *int                            `json:"-" url:"attempts,omitempty"`
	ResponseStatus *int                            `json:"-" url:"response_status,omitempty"`
	SuccessfulAt   *time.Time                      `json:"-" url:"successful_at,omitempty"`
	CreatedAt      *time.Time                      `json:"-" url:"created_at,omitempty"`
	ErrorCode      *AttemptErrorCodes              `json:"-" url:"error_code,omitempty"`
	CliId          *string                         `json:"-" url:"cli_id,omitempty"`
	LastAttemptAt  *time.Time                      `json:"-" url:"last_attempt_at,omitempty"`
	SearchTerm     *string                         `json:"-" url:"search_term,omitempty"`
	Headers        *string                         `json:"-" url:"headers,omitempty"`
	Body           *string                         `json:"-" url:"body,omitempty"`
	ParsedQuery    *string                         `json:"-" url:"parsed_query,omitempty"`
	Path           *string                         `json:"-" url:"path,omitempty"`
	CliUserId      []*string                       `json:"-" url:"cli_user_id,omitempty"`
	IssueId        []*string                       `json:"-" url:"issue_id,omitempty"`
	EventDataId    []*string                       `json:"-" url:"event_data_id,omitempty"`
	BulkRetryId    []*string                       `json:"-" url:"bulk_retry_id,omitempty"`
	Include        *string                         `json:"-" url:"include,omitempty"`
	OrderBy        *RequestListEventRequestOrderBy `json:"-" url:"order_by,omitempty"`
	Dir            *RequestListEventRequestDir     `json:"-" url:"dir,omitempty"`
	Limit          *int                            `json:"-" url:"limit,omitempty"`
	Next           *string                         `json:"-" url:"next,omitempty"`
	Prev           *string                         `json:"-" url:"prev,omitempty"`
}

type RequestListIgnoredEventRequest struct {
	Id      []*string                              `json:"-" url:"id,omitempty"`
	OrderBy *RequestListIgnoredEventRequestOrderBy `json:"-" url:"order_by,omitempty"`
	Dir     *RequestListIgnoredEventRequestDir     `json:"-" url:"dir,omitempty"`
	Limit   *int                                   `json:"-" url:"limit,omitempty"`
	Next    *string                                `json:"-" url:"next,omitempty"`
	Prev    *string                                `json:"-" url:"prev,omitempty"`
}

type RequestRetryRequest struct {
	// Subset of webhook_ids to re-run the event logic on. Useful to retry only specific ignored_events
	WebhookIds []string `json:"webhook_ids,omitempty" url:"-"`
}

// Sort direction
type RequestListEventRequestDir string

const (
	RequestListEventRequestDirAsc  RequestListEventRequestDir = "asc"
	RequestListEventRequestDirDesc RequestListEventRequestDir = "desc"
)

func NewRequestListEventRequestDirFromString(s string) (RequestListEventRequestDir, error) {
	switch s {
	case "asc":
		return RequestListEventRequestDirAsc, nil
	case "desc":
		return RequestListEventRequestDirDesc, nil
	}
	var t RequestListEventRequestDir
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (r RequestListEventRequestDir) Ptr() *RequestListEventRequestDir {
	return &r
}

// Sort key
type RequestListEventRequestOrderBy string

const (
	RequestListEventRequestOrderByLastAttemptAt RequestListEventRequestOrderBy = "last_attempt_at"
	RequestListEventRequestOrderByCreatedAt     RequestListEventRequestOrderBy = "created_at"
)

func NewRequestListEventRequestOrderByFromString(s string) (RequestListEventRequestOrderBy, error) {
	switch s {
	case "last_attempt_at":
		return RequestListEventRequestOrderByLastAttemptAt, nil
	case "created_at":
		return RequestListEventRequestOrderByCreatedAt, nil
	}
	var t RequestListEventRequestOrderBy
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (r RequestListEventRequestOrderBy) Ptr() *RequestListEventRequestOrderBy {
	return &r
}

type RequestListIgnoredEventRequestDir string

const (
	RequestListIgnoredEventRequestDirAsc  RequestListIgnoredEventRequestDir = "asc"
	RequestListIgnoredEventRequestDirDesc RequestListIgnoredEventRequestDir = "desc"
)

func NewRequestListIgnoredEventRequestDirFromString(s string) (RequestListIgnoredEventRequestDir, error) {
	switch s {
	case "asc":
		return RequestListIgnoredEventRequestDirAsc, nil
	case "desc":
		return RequestListIgnoredEventRequestDirDesc, nil
	}
	var t RequestListIgnoredEventRequestDir
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (r RequestListIgnoredEventRequestDir) Ptr() *RequestListIgnoredEventRequestDir {
	return &r
}

type RequestListIgnoredEventRequestOrderBy string

const (
	RequestListIgnoredEventRequestOrderByCreatedAt RequestListIgnoredEventRequestOrderBy = "created_at"
)

func NewRequestListIgnoredEventRequestOrderByFromString(s string) (RequestListIgnoredEventRequestOrderBy, error) {
	switch s {
	case "created_at":
		return RequestListIgnoredEventRequestOrderByCreatedAt, nil
	}
	var t RequestListIgnoredEventRequestOrderBy
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (r RequestListIgnoredEventRequestOrderBy) Ptr() *RequestListIgnoredEventRequestOrderBy {
	return &r
}

// Sort direction
type RequestListRequestDir string

const (
	RequestListRequestDirAsc  RequestListRequestDir = "asc"
	RequestListRequestDirDesc RequestListRequestDir = "desc"
)

func NewRequestListRequestDirFromString(s string) (RequestListRequestDir, error) {
	switch s {
	case "asc":
		return RequestListRequestDirAsc, nil
	case "desc":
		return RequestListRequestDirDesc, nil
	}
	var t RequestListRequestDir
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (r RequestListRequestDir) Ptr() *RequestListRequestDir {
	return &r
}

// Sort key
type RequestListRequestOrderBy string

const (
	RequestListRequestOrderByIngestedAt RequestListRequestOrderBy = "ingested_at"
	RequestListRequestOrderByCreatedAt  RequestListRequestOrderBy = "created_at"
)

func NewRequestListRequestOrderByFromString(s string) (RequestListRequestOrderBy, error) {
	switch s {
	case "ingested_at":
		return RequestListRequestOrderByIngestedAt, nil
	case "created_at":
		return RequestListRequestOrderByCreatedAt, nil
	}
	var t RequestListRequestOrderBy
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (r RequestListRequestOrderBy) Ptr() *RequestListRequestOrderBy {
	return &r
}

// Filter by status
type RequestListRequestStatus string

const (
	RequestListRequestStatusAccepted RequestListRequestStatus = "accepted"
	RequestListRequestStatusRejected RequestListRequestStatus = "rejected"
)

func NewRequestListRequestStatusFromString(s string) (RequestListRequestStatus, error) {
	switch s {
	case "accepted":
		return RequestListRequestStatusAccepted, nil
	case "rejected":
		return RequestListRequestStatusRejected, nil
	}
	var t RequestListRequestStatus
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (r RequestListRequestStatus) Ptr() *RequestListRequestStatus {
	return &r
}
