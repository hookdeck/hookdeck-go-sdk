// This file was auto-generated by Fern from our API Definition.

package api

import (
	core "github.com/hookdeck/hookdeck-go-sdk/core"
	time "time"
)

type SourceCreateRequest struct {
	// A unique name for the source <span style="white-space: nowrap">`<= 155 characters`</span>
	Name string `json:"name"`
	// Description for the source
	Description        *core.Optional[string]                  `json:"description,omitempty"`
	AllowedHttpMethods *core.Optional[SourceAllowedHttpMethod] `json:"allowed_http_methods,omitempty"`
	CustomResponse     *core.Optional[SourceCustomResponse]    `json:"custom_response,omitempty"`
	Verification       *core.Optional[VerificationConfig]      `json:"verification,omitempty"`
}

type SourceListRequest struct {
	Id         *string               `json:"-"`
	Name       *string               `json:"-"`
	Archived   *bool                 `json:"-"`
	ArchivedAt *time.Time            `json:"-"`
	OrderBy    *string               `json:"-"`
	Dir        *SourceListRequestDir `json:"-"`
	Limit      *int                  `json:"-"`
	Next       *string               `json:"-"`
	Prev       *string               `json:"-"`
}

type SourceRetrieveRequest struct {
	Include *string `json:"-"`
}

type SourceUpdateRequest struct {
	// A unique name for the source <span style="white-space: nowrap">`<= 155 characters`</span>
	Name *core.Optional[string] `json:"name,omitempty"`
	// Description for the source
	Description        *core.Optional[string]                  `json:"description,omitempty"`
	AllowedHttpMethods *core.Optional[SourceAllowedHttpMethod] `json:"allowed_http_methods,omitempty"`
	CustomResponse     *core.Optional[SourceCustomResponse]    `json:"custom_response,omitempty"`
	Verification       *core.Optional[VerificationConfig]      `json:"verification,omitempty"`
}

type SourceUpsertRequest struct {
	// A unique name for the source <span style="white-space: nowrap">`<= 155 characters`</span>
	Name string `json:"name"`
	// Description for the source
	Description        *core.Optional[string]                  `json:"description,omitempty"`
	AllowedHttpMethods *core.Optional[SourceAllowedHttpMethod] `json:"allowed_http_methods,omitempty"`
	CustomResponse     *core.Optional[SourceCustomResponse]    `json:"custom_response,omitempty"`
	Verification       *core.Optional[VerificationConfig]      `json:"verification,omitempty"`
}
