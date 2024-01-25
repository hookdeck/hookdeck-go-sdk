// This file was auto-generated by Fern from our API Definition.

package api

import (
	fmt "fmt"
	time "time"
)

type EventListRequest struct {
	Id             *string                  `json:"-"`
	Status         *EventStatus             `json:"-"`
	WebhookId      *string                  `json:"-"`
	DestinationId  *string                  `json:"-"`
	SourceId       *string                  `json:"-"`
	Attempts       *int                     `json:"-"`
	ResponseStatus *int                     `json:"-"`
	SuccessfulAt   *time.Time               `json:"-"`
	CreatedAt      *time.Time               `json:"-"`
	ErrorCode      *AttemptErrorCodes       `json:"-"`
	CliId          *string                  `json:"-"`
	LastAttemptAt  *time.Time               `json:"-"`
	SearchTerm     *string                  `json:"-"`
	Headers        *string                  `json:"-"`
	Body           *string                  `json:"-"`
	ParsedQuery    *string                  `json:"-"`
	Path           *string                  `json:"-"`
	CliUserId      *string                  `json:"-"`
	IssueId        *string                  `json:"-"`
	EventDataId    *string                  `json:"-"`
	BulkRetryId    *string                  `json:"-"`
	Include        *string                  `json:"-"`
	OrderBy        *EventListRequestOrderBy `json:"-"`
	Dir            *EventListRequestDir     `json:"-"`
	Limit          *int                     `json:"-"`
	Next           *string                  `json:"-"`
	Prev           *string                  `json:"-"`
}

// Sort direction
type EventListRequestDir string

const (
	EventListRequestDirAsc  EventListRequestDir = "asc"
	EventListRequestDirDesc EventListRequestDir = "desc"
)

func NewEventListRequestDirFromString(s string) (EventListRequestDir, error) {
	switch s {
	case "asc":
		return EventListRequestDirAsc, nil
	case "desc":
		return EventListRequestDirDesc, nil
	}
	var t EventListRequestDir
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e EventListRequestDir) Ptr() *EventListRequestDir {
	return &e
}

// Sort key
type EventListRequestOrderBy string

const (
	EventListRequestOrderByLastAttemptAt EventListRequestOrderBy = "last_attempt_at"
	EventListRequestOrderByCreatedAt     EventListRequestOrderBy = "created_at"
)

func NewEventListRequestOrderByFromString(s string) (EventListRequestOrderBy, error) {
	switch s {
	case "last_attempt_at":
		return EventListRequestOrderByLastAttemptAt, nil
	case "created_at":
		return EventListRequestOrderByCreatedAt, nil
	}
	var t EventListRequestOrderBy
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e EventListRequestOrderBy) Ptr() *EventListRequestOrderBy {
	return &e
}