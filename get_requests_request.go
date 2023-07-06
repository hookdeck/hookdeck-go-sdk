// This file was auto-generated by Fern from our API Definition.

package api

// GetRequestsRequest is an in-lined request used by the GetRequests endpoint.
type GetRequestsRequest struct {
	Id             *string                                             `json:"-"`
	Status         *GetRequestsRequestStatus                           `json:"-"`
	RejectionCause *RequestRejectionCause                              `json:"-"`
	IgnoredCount   *int                                                `json:"-"`
	EventsCount    *int                                                `json:"-"`
	SourceId       *string                                             `json:"-"`
	Verified       *bool                                               `json:"-"`
	Headers        *string                                             `json:"-"`
	Body           *string                                             `json:"-"`
	ParsedQuery    *string                                             `json:"-"`
	Path           *string                                             `json:"-"`
	CreatedAt      *string                                             `json:"-"`
	IngestedAt     *string                                             `json:"-"`
	BulkRetryId    *string                                             `json:"-"`
	Include        *string                                             `json:"-"`
	OrderBy        *GetRequestsRequestOrderByGetRequestsRequestOrderBy `json:"-"`
	Dir            *GetRequestsRequestDirGetRequestsRequestDir         `json:"-"`
	Limit          *int                                                `json:"-"`
	Next           *string                                             `json:"-"`
	Prev           *string                                             `json:"-"`
}
