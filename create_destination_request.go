// This file was auto-generated by Fern from our API Definition.

package api

// CreateDestinationRequest is an in-lined request used by the CreateDestination endpoint.
type CreateDestinationRequest struct {
	// Name for the destination <span style="white-space: nowrap">`<= 155 characters`</span>
	Name string `json:"name,omitempty"`
	// Endpoint of the destination
	Url *string `json:"url,omitempty"`
	// Path for the CLI destination
	CliPath *string `json:"cli_path,omitempty"`
	// Period to rate limit attempts
	RateLimitPeriod *CreateDestinationRequestRateLimitPeriod `json:"rate_limit_period,omitempty"`
	// Limit event attempts to receive per period
	RateLimit              *int  `json:"rate_limit,omitempty"`
	PathForwardingDisabled *bool `json:"path_forwarding_disabled,omitempty"`
}
