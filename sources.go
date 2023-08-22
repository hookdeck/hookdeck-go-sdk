// This file was auto-generated by Fern from our API Definition.

package api

import (
	time "time"

	core "github.com/hookdeck/hookdeck-go-sdk/core"
)

type CreateSourceRequest struct {
	// A unique name for the source <span style="white-space: nowrap">`<= 155 characters`</span>
	Name string `json:"name"`
	// Description for the source
	Description        *string                  `json:"description,omitempty"`
	AllowedHttpMethods *SourceAllowedHttpMethod `json:"allowed_http_methods,omitempty"`
	CustomResponse     *SourceCustomResponse    `json:"custom_response,omitempty"`
	Verification       *VerificationConfig      `json:"verification,omitempty"`
}

type GetSourceRequest struct {
	Include *string `json:"-"`
}

type GetSourcesRequest struct {
	Id         *string               `json:"-"`
	Name       *string               `json:"-"`
	Archived   *bool                 `json:"-"`
	ArchivedAt *time.Time            `json:"-"`
	OrderBy    *string               `json:"-"`
	Dir        *GetSourcesRequestDir `json:"-"`
	Limit      *int                  `json:"-"`
	Next       *string               `json:"-"`
	Prev       *string               `json:"-"`
}

type UpdateSourceRequest struct {
	// A unique name for the source <span style="white-space: nowrap">`<= 155 characters`</span>
	Name *core.Optional[string] `json:"name,omitempty"`
	// Description for the source
	Description        *core.Optional[string]                  `json:"description,omitempty"`
	AllowedHttpMethods *core.Optional[SourceAllowedHttpMethod] `json:"allowed_http_methods,omitempty"`
	CustomResponse     *core.Optional[SourceCustomResponse]    `json:"custom_response,omitempty"`
	Verification       *core.Optional[VerificationConfig]      `json:"verification,omitempty"`
}

type UpsertSourceRequest struct {
	// A unique name for the source <span style="white-space: nowrap">`<= 155 characters`</span>
	Name string `json:"name"`
	// Description for the source
	Description        *string                  `json:"description,omitempty"`
	AllowedHttpMethods *SourceAllowedHttpMethod `json:"allowed_http_methods,omitempty"`
	CustomResponse     *SourceCustomResponse    `json:"custom_response,omitempty"`
	Verification       *VerificationConfig      `json:"verification,omitempty"`
}
