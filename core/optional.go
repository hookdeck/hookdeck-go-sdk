package core

import (
	"encoding/json"
	"fmt"
	"time"
)

// Optional is a wrapper used to distinguish zero values from
// null or omitted fields.
//
// To instantiate an Optional, use the `Optional()` and `Null()`
// helpers exported from the root package.
type Optional[T any] struct {
	Value T
	Null  bool
}

func (o *Optional[T]) String() string {
	if o == nil {
		return ""
	}
	if s, ok := any(o.Value).(fmt.Stringer); ok {
		return s.String()
	}
	return fmt.Sprintf("%#v", o.Value)
}

func (o *Optional[T]) MarshalJSON() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	if o.Null {
		return []byte("null"), nil
	}
	return json.Marshal(&o.Value)
}

// NewDateFromOptional returns a new *DateTime from the given optional.
func NewDateFromOptional(optional *Optional[time.Time]) *Date {
	if optional == nil {
		return nil
	}
	return &Date{t: &optional.Value}
}

// NewDateTimeFromOptional returns a new *DateTime from the given optional.
func NewDateTimeFromOptional(optional *Optional[time.Time]) *DateTime {
	if optional == nil {
		return nil
	}
	return &DateTime{t: &optional.Value}
}
