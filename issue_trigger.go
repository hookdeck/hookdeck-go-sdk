// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
	core "github.com/hookdeck/hookdeck-go-sdk/core"
	time "time"
)

type IssueTriggerCreateRequest struct {
	Type IssueType `json:"type" url:"-"`
	// Configuration object for the specific issue type selected
	Configs  *core.Optional[IssueTriggerCreateRequestConfigs] `json:"configs,omitempty" url:"-"`
	Channels *core.Optional[IssueTriggerChannels]             `json:"channels,omitempty" url:"-"`
	// Optional unique name to use as reference when using the API
	Name *core.Optional[string] `json:"name,omitempty" url:"-"`
}

type IssueTriggerListRequest struct {
	Name       *string                         `json:"-" url:"name,omitempty"`
	Type       *IssueType                      `json:"-" url:"type,omitempty"`
	DisabledAt *time.Time                      `json:"-" url:"disabled_at,omitempty"`
	OrderBy    *IssueTriggerListRequestOrderBy `json:"-" url:"order_by,omitempty"`
	Dir        *IssueTriggerListRequestDir     `json:"-" url:"dir,omitempty"`
	Limit      *int                            `json:"-" url:"limit,omitempty"`
	Next       *string                         `json:"-" url:"next,omitempty"`
	Prev       *string                         `json:"-" url:"prev,omitempty"`
}

// Configuration object for the specific issue type selected
type IssueTriggerCreateRequestConfigs struct {
	IssueTriggerDeliveryConfigs       *IssueTriggerDeliveryConfigs
	IssueTriggerTransformationConfigs *IssueTriggerTransformationConfigs
	IssueTriggerBackpressureConfigs   *IssueTriggerBackpressureConfigs
}

func NewIssueTriggerCreateRequestConfigsFromIssueTriggerDeliveryConfigs(value *IssueTriggerDeliveryConfigs) *IssueTriggerCreateRequestConfigs {
	return &IssueTriggerCreateRequestConfigs{IssueTriggerDeliveryConfigs: value}
}

func NewIssueTriggerCreateRequestConfigsFromIssueTriggerTransformationConfigs(value *IssueTriggerTransformationConfigs) *IssueTriggerCreateRequestConfigs {
	return &IssueTriggerCreateRequestConfigs{IssueTriggerTransformationConfigs: value}
}

func NewIssueTriggerCreateRequestConfigsFromIssueTriggerBackpressureConfigs(value *IssueTriggerBackpressureConfigs) *IssueTriggerCreateRequestConfigs {
	return &IssueTriggerCreateRequestConfigs{IssueTriggerBackpressureConfigs: value}
}

func (i *IssueTriggerCreateRequestConfigs) UnmarshalJSON(data []byte) error {
	valueIssueTriggerDeliveryConfigs := new(IssueTriggerDeliveryConfigs)
	if err := json.Unmarshal(data, &valueIssueTriggerDeliveryConfigs); err == nil {
		i.IssueTriggerDeliveryConfigs = valueIssueTriggerDeliveryConfigs
		return nil
	}
	valueIssueTriggerTransformationConfigs := new(IssueTriggerTransformationConfigs)
	if err := json.Unmarshal(data, &valueIssueTriggerTransformationConfigs); err == nil {
		i.IssueTriggerTransformationConfigs = valueIssueTriggerTransformationConfigs
		return nil
	}
	valueIssueTriggerBackpressureConfigs := new(IssueTriggerBackpressureConfigs)
	if err := json.Unmarshal(data, &valueIssueTriggerBackpressureConfigs); err == nil {
		i.IssueTriggerBackpressureConfigs = valueIssueTriggerBackpressureConfigs
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, i)
}

func (i IssueTriggerCreateRequestConfigs) MarshalJSON() ([]byte, error) {
	if i.IssueTriggerDeliveryConfigs != nil {
		return json.Marshal(i.IssueTriggerDeliveryConfigs)
	}
	if i.IssueTriggerTransformationConfigs != nil {
		return json.Marshal(i.IssueTriggerTransformationConfigs)
	}
	if i.IssueTriggerBackpressureConfigs != nil {
		return json.Marshal(i.IssueTriggerBackpressureConfigs)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", i)
}

type IssueTriggerCreateRequestConfigsVisitor interface {
	VisitIssueTriggerDeliveryConfigs(*IssueTriggerDeliveryConfigs) error
	VisitIssueTriggerTransformationConfigs(*IssueTriggerTransformationConfigs) error
	VisitIssueTriggerBackpressureConfigs(*IssueTriggerBackpressureConfigs) error
}

func (i *IssueTriggerCreateRequestConfigs) Accept(visitor IssueTriggerCreateRequestConfigsVisitor) error {
	if i.IssueTriggerDeliveryConfigs != nil {
		return visitor.VisitIssueTriggerDeliveryConfigs(i.IssueTriggerDeliveryConfigs)
	}
	if i.IssueTriggerTransformationConfigs != nil {
		return visitor.VisitIssueTriggerTransformationConfigs(i.IssueTriggerTransformationConfigs)
	}
	if i.IssueTriggerBackpressureConfigs != nil {
		return visitor.VisitIssueTriggerBackpressureConfigs(i.IssueTriggerBackpressureConfigs)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", i)
}

type IssueTriggerListRequestDir string

const (
	IssueTriggerListRequestDirAsc  IssueTriggerListRequestDir = "asc"
	IssueTriggerListRequestDirDesc IssueTriggerListRequestDir = "desc"
)

func NewIssueTriggerListRequestDirFromString(s string) (IssueTriggerListRequestDir, error) {
	switch s {
	case "asc":
		return IssueTriggerListRequestDirAsc, nil
	case "desc":
		return IssueTriggerListRequestDirDesc, nil
	}
	var t IssueTriggerListRequestDir
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (i IssueTriggerListRequestDir) Ptr() *IssueTriggerListRequestDir {
	return &i
}

type IssueTriggerListRequestOrderBy string

const (
	IssueTriggerListRequestOrderByCreatedAt IssueTriggerListRequestOrderBy = "created_at"
	IssueTriggerListRequestOrderByType      IssueTriggerListRequestOrderBy = "type"
)

func NewIssueTriggerListRequestOrderByFromString(s string) (IssueTriggerListRequestOrderBy, error) {
	switch s {
	case "created_at":
		return IssueTriggerListRequestOrderByCreatedAt, nil
	case "type":
		return IssueTriggerListRequestOrderByType, nil
	}
	var t IssueTriggerListRequestOrderBy
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (i IssueTriggerListRequestOrderBy) Ptr() *IssueTriggerListRequestOrderBy {
	return &i
}

// Configuration object for the specific issue type selected
type IssueTriggerUpdateRequestConfigs struct {
	IssueTriggerDeliveryConfigs       *IssueTriggerDeliveryConfigs
	IssueTriggerTransformationConfigs *IssueTriggerTransformationConfigs
	IssueTriggerBackpressureConfigs   *IssueTriggerBackpressureConfigs
}

func NewIssueTriggerUpdateRequestConfigsFromIssueTriggerDeliveryConfigs(value *IssueTriggerDeliveryConfigs) *IssueTriggerUpdateRequestConfigs {
	return &IssueTriggerUpdateRequestConfigs{IssueTriggerDeliveryConfigs: value}
}

func NewIssueTriggerUpdateRequestConfigsFromIssueTriggerTransformationConfigs(value *IssueTriggerTransformationConfigs) *IssueTriggerUpdateRequestConfigs {
	return &IssueTriggerUpdateRequestConfigs{IssueTriggerTransformationConfigs: value}
}

func NewIssueTriggerUpdateRequestConfigsFromIssueTriggerBackpressureConfigs(value *IssueTriggerBackpressureConfigs) *IssueTriggerUpdateRequestConfigs {
	return &IssueTriggerUpdateRequestConfigs{IssueTriggerBackpressureConfigs: value}
}

func (i *IssueTriggerUpdateRequestConfigs) UnmarshalJSON(data []byte) error {
	valueIssueTriggerDeliveryConfigs := new(IssueTriggerDeliveryConfigs)
	if err := json.Unmarshal(data, &valueIssueTriggerDeliveryConfigs); err == nil {
		i.IssueTriggerDeliveryConfigs = valueIssueTriggerDeliveryConfigs
		return nil
	}
	valueIssueTriggerTransformationConfigs := new(IssueTriggerTransformationConfigs)
	if err := json.Unmarshal(data, &valueIssueTriggerTransformationConfigs); err == nil {
		i.IssueTriggerTransformationConfigs = valueIssueTriggerTransformationConfigs
		return nil
	}
	valueIssueTriggerBackpressureConfigs := new(IssueTriggerBackpressureConfigs)
	if err := json.Unmarshal(data, &valueIssueTriggerBackpressureConfigs); err == nil {
		i.IssueTriggerBackpressureConfigs = valueIssueTriggerBackpressureConfigs
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, i)
}

func (i IssueTriggerUpdateRequestConfigs) MarshalJSON() ([]byte, error) {
	if i.IssueTriggerDeliveryConfigs != nil {
		return json.Marshal(i.IssueTriggerDeliveryConfigs)
	}
	if i.IssueTriggerTransformationConfigs != nil {
		return json.Marshal(i.IssueTriggerTransformationConfigs)
	}
	if i.IssueTriggerBackpressureConfigs != nil {
		return json.Marshal(i.IssueTriggerBackpressureConfigs)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", i)
}

type IssueTriggerUpdateRequestConfigsVisitor interface {
	VisitIssueTriggerDeliveryConfigs(*IssueTriggerDeliveryConfigs) error
	VisitIssueTriggerTransformationConfigs(*IssueTriggerTransformationConfigs) error
	VisitIssueTriggerBackpressureConfigs(*IssueTriggerBackpressureConfigs) error
}

func (i *IssueTriggerUpdateRequestConfigs) Accept(visitor IssueTriggerUpdateRequestConfigsVisitor) error {
	if i.IssueTriggerDeliveryConfigs != nil {
		return visitor.VisitIssueTriggerDeliveryConfigs(i.IssueTriggerDeliveryConfigs)
	}
	if i.IssueTriggerTransformationConfigs != nil {
		return visitor.VisitIssueTriggerTransformationConfigs(i.IssueTriggerTransformationConfigs)
	}
	if i.IssueTriggerBackpressureConfigs != nil {
		return visitor.VisitIssueTriggerBackpressureConfigs(i.IssueTriggerBackpressureConfigs)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", i)
}

// Configuration object for the specific issue type selected
type IssueTriggerUpsertRequestConfigs struct {
	IssueTriggerDeliveryConfigs       *IssueTriggerDeliveryConfigs
	IssueTriggerTransformationConfigs *IssueTriggerTransformationConfigs
	IssueTriggerBackpressureConfigs   *IssueTriggerBackpressureConfigs
}

func NewIssueTriggerUpsertRequestConfigsFromIssueTriggerDeliveryConfigs(value *IssueTriggerDeliveryConfigs) *IssueTriggerUpsertRequestConfigs {
	return &IssueTriggerUpsertRequestConfigs{IssueTriggerDeliveryConfigs: value}
}

func NewIssueTriggerUpsertRequestConfigsFromIssueTriggerTransformationConfigs(value *IssueTriggerTransformationConfigs) *IssueTriggerUpsertRequestConfigs {
	return &IssueTriggerUpsertRequestConfigs{IssueTriggerTransformationConfigs: value}
}

func NewIssueTriggerUpsertRequestConfigsFromIssueTriggerBackpressureConfigs(value *IssueTriggerBackpressureConfigs) *IssueTriggerUpsertRequestConfigs {
	return &IssueTriggerUpsertRequestConfigs{IssueTriggerBackpressureConfigs: value}
}

func (i *IssueTriggerUpsertRequestConfigs) UnmarshalJSON(data []byte) error {
	valueIssueTriggerDeliveryConfigs := new(IssueTriggerDeliveryConfigs)
	if err := json.Unmarshal(data, &valueIssueTriggerDeliveryConfigs); err == nil {
		i.IssueTriggerDeliveryConfigs = valueIssueTriggerDeliveryConfigs
		return nil
	}
	valueIssueTriggerTransformationConfigs := new(IssueTriggerTransformationConfigs)
	if err := json.Unmarshal(data, &valueIssueTriggerTransformationConfigs); err == nil {
		i.IssueTriggerTransformationConfigs = valueIssueTriggerTransformationConfigs
		return nil
	}
	valueIssueTriggerBackpressureConfigs := new(IssueTriggerBackpressureConfigs)
	if err := json.Unmarshal(data, &valueIssueTriggerBackpressureConfigs); err == nil {
		i.IssueTriggerBackpressureConfigs = valueIssueTriggerBackpressureConfigs
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, i)
}

func (i IssueTriggerUpsertRequestConfigs) MarshalJSON() ([]byte, error) {
	if i.IssueTriggerDeliveryConfigs != nil {
		return json.Marshal(i.IssueTriggerDeliveryConfigs)
	}
	if i.IssueTriggerTransformationConfigs != nil {
		return json.Marshal(i.IssueTriggerTransformationConfigs)
	}
	if i.IssueTriggerBackpressureConfigs != nil {
		return json.Marshal(i.IssueTriggerBackpressureConfigs)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", i)
}

type IssueTriggerUpsertRequestConfigsVisitor interface {
	VisitIssueTriggerDeliveryConfigs(*IssueTriggerDeliveryConfigs) error
	VisitIssueTriggerTransformationConfigs(*IssueTriggerTransformationConfigs) error
	VisitIssueTriggerBackpressureConfigs(*IssueTriggerBackpressureConfigs) error
}

func (i *IssueTriggerUpsertRequestConfigs) Accept(visitor IssueTriggerUpsertRequestConfigsVisitor) error {
	if i.IssueTriggerDeliveryConfigs != nil {
		return visitor.VisitIssueTriggerDeliveryConfigs(i.IssueTriggerDeliveryConfigs)
	}
	if i.IssueTriggerTransformationConfigs != nil {
		return visitor.VisitIssueTriggerTransformationConfigs(i.IssueTriggerTransformationConfigs)
	}
	if i.IssueTriggerBackpressureConfigs != nil {
		return visitor.VisitIssueTriggerBackpressureConfigs(i.IssueTriggerBackpressureConfigs)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", i)
}

type IssueTriggerUpdateRequest struct {
	// Configuration object for the specific issue type selected
	Configs  *core.Optional[IssueTriggerUpdateRequestConfigs] `json:"configs,omitempty" url:"-"`
	Channels *core.Optional[IssueTriggerChannels]             `json:"channels,omitempty" url:"-"`
	// Date when the issue trigger was disabled
	DisabledAt *core.Optional[time.Time] `json:"disabled_at,omitempty" url:"-"`
	// Optional unique name to use as reference when using the API
	Name *core.Optional[string] `json:"name,omitempty" url:"-"`
}

func (i *IssueTriggerUpdateRequest) UnmarshalJSON(data []byte) error {
	type unmarshaler IssueTriggerUpdateRequest
	var body unmarshaler
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	*i = IssueTriggerUpdateRequest(body)
	return nil
}

func (i *IssueTriggerUpdateRequest) MarshalJSON() ([]byte, error) {
	type embed IssueTriggerUpdateRequest
	var marshaler = struct {
		embed
		DisabledAt *core.DateTime `json:"disabled_at,omitempty"`
	}{
		embed:      embed(*i),
		DisabledAt: core.NewOptionalDateTime(i.DisabledAt),
	}
	return json.Marshal(marshaler)
}

type IssueTriggerUpsertRequest struct {
	Type IssueType `json:"type" url:"-"`
	// Configuration object for the specific issue type selected
	Configs  *core.Optional[IssueTriggerUpsertRequestConfigs] `json:"configs,omitempty" url:"-"`
	Channels *core.Optional[IssueTriggerChannels]             `json:"channels,omitempty" url:"-"`
	// Required unique name to use as reference when using the API
	Name string `json:"name" url:"-"`
}
