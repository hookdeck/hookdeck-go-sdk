package api

import core "github.com/hookdeck/hookdeck-go-sdk/core"

// O initializes an optional field.
func O[T any](value T) *core.Optional[T] {
	return &core.Optional[T]{
		Value: value,
	}
}

// Null initializes an optional field that will be sent as
// an explicit null value.
func Null[T any]() *core.Optional[T] {
	return &core.Optional[T]{
		Null: true,
	}
}
