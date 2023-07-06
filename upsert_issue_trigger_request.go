// This file was auto-generated by Fern from our API Definition.

package api

// UpsertIssueTriggerRequest is an in-lined request used by the UpsertIssueTrigger endpoint.
type UpsertIssueTriggerRequest struct {
	Type IssueType `json:"type,omitempty"`
	// Configuration object for the specific issue type selected
	Configs  *UpsertIssueTriggerRequestConfigs `json:"configs,omitempty"`
	Channels *IssueTriggerChannels             `json:"channels,omitempty"`
	// Required unique name to use as reference when using the API <span style="white-space: nowrap">`<= 255 characters`</span>
	Name string `json:"name,omitempty"`
}
