// This file was auto-generated by Fern from our API Definition.

package api

// UpdateConnectionRequest is an in-lined request used by the UpdateConnection endpoint.
type UpdateConnectionRequest struct {
	// Unique name of the connection for the source <span style="white-space: nowrap">`<= 155 characters`</span>
	Name *string `json:"name,omitempty"`
	// Ruleset input object
	Ruleset *UpdateConnectionRequestRuleset `json:"ruleset,omitempty"`
	// ID of a rule to bind to the connection. Default to the Workspace default ruleset
	RulesetId *string `json:"ruleset_id,omitempty"`
	// Array of rules to apply
	Rules *[]*Rule `json:"rules,omitempty"`
}
