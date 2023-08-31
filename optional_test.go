package api

import (
	"testing"

	core "github.com/hookdeck/hookdeck-go-sdk/core"
	"github.com/stretchr/testify/assert"
)

func TestOptionalOrNull(t *testing.T) {
	t.Run("nil primitive", func(t *testing.T) {
		var value *string
		optional := OptionalOrNull(value)
		assert.Equal(t, &core.Optional[string]{Null: true}, optional)
	})

	t.Run("zero primitive", func(t *testing.T) {
		var zero int
		optional := OptionalOrNull(&zero)
		assert.Equal(t, &core.Optional[int]{Value: zero}, optional)
	})

	t.Run("non-zero primitive", func(t *testing.T) {
		value := "test"
		optional := OptionalOrNull(&value)
		assert.Equal(t, &core.Optional[string]{Value: value}, optional)
	})

	t.Run("nil type", func(t *testing.T) {
		type foo struct {
			Name string
		}
		var zero *foo
		optional := OptionalOrNull(zero)
		assert.Equal(t, &core.Optional[foo]{Null: true}, optional)
	})

	t.Run("zero type", func(t *testing.T) {
		type foo struct {
			Name string
		}
		zero := new(foo)
		optional := OptionalOrNull(zero)
		assert.Equal(t, &core.Optional[foo]{Value: *zero}, optional)
	})

	t.Run("non-zero type", func(t *testing.T) {
		type foo struct {
			Name string
		}
		value := &foo{Name: "test"}
		optional := OptionalOrNull(value)
		assert.Equal(t, &core.Optional[foo]{Value: *value}, optional)
	})
}
