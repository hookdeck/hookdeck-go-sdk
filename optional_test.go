package api

import (
	"testing"

	core "github.com/hookdeck/hookdeck-go-sdk/core"
	"github.com/stretchr/testify/assert"
)

func TestOptionalOrNull(t *testing.T) {
	t.Run("nil map", func(t *testing.T) {
		var value map[string]string
		optional := OptionalOrNull(value)
		assert.Equal(t, &core.Optional[map[string]string]{Null: true}, optional)
	})

	t.Run("zero map", func(t *testing.T) {
		value := make(map[string]string)
		optional := OptionalOrNull(value)
		assert.Equal(t, &core.Optional[map[string]string]{Value: value}, optional)
	})

	t.Run("non-zero map", func(t *testing.T) {
		value := map[string]string{
			"foo": "bar",
		}
		optional := OptionalOrNull(value)
		assert.Equal(t, &core.Optional[map[string]string]{Value: value}, optional)
	})

	t.Run("nil slice", func(t *testing.T) {
		var value []string
		optional := OptionalOrNull(value)
		assert.Equal(t, &core.Optional[[]string]{Null: true}, optional)
	})

	t.Run("zero slice", func(t *testing.T) {
		value := make([]string, 0, 0)
		optional := OptionalOrNull(value)
		assert.Equal(t, &core.Optional[[]string]{Value: value}, optional)
	})

	t.Run("non-zero slice", func(t *testing.T) {
		value := []string{
			"foo",
			"bar",
		}
		optional := OptionalOrNull(value)
		assert.Equal(t, &core.Optional[[]string]{Value: value}, optional)
	})

	t.Run("zero type", func(t *testing.T) {
		type foo struct {
			Name string
		}
		zero := foo{}
		optional := OptionalOrNull(zero)
		assert.Equal(t, &core.Optional[foo]{Value: zero}, optional)
	})

	t.Run("non-zero type", func(t *testing.T) {
		type foo struct {
			Name string
		}
		value := foo{Name: "test"}
		optional := OptionalOrNull(value)
		assert.Equal(t, &core.Optional[foo]{Value: value}, optional)
	})

	t.Run("zero interface", func(t *testing.T) {
		var s string
		var iface any = s
		optional := OptionalOrNull(iface)
		assert.Equal(t, &core.Optional[any]{Value: iface}, optional)
	})

	t.Run("non-zero interface", func(t *testing.T) {
		s := "foo"
		var iface any = s
		optional := OptionalOrNull(iface)
		assert.Equal(t, &core.Optional[any]{Value: iface}, optional)
	})
}

func TestOptionalOrNullPtr(t *testing.T) {
	t.Run("nil primitive", func(t *testing.T) {
		var value *string
		optional := OptionalOrNullPtr(value)
		assert.Equal(t, &core.Optional[string]{Null: true}, optional)
	})

	t.Run("zero primitive", func(t *testing.T) {
		var zero int
		optional := OptionalOrNullPtr(&zero)
		assert.Equal(t, &core.Optional[int]{Value: zero}, optional)
	})

	t.Run("non-zero primitive", func(t *testing.T) {
		value := "test"
		optional := OptionalOrNullPtr(&value)
		assert.Equal(t, &core.Optional[string]{Value: value}, optional)
	})

	t.Run("nil type", func(t *testing.T) {
		type foo struct {
			Name string
		}
		var zero *foo
		optional := OptionalOrNullPtr(zero)
		assert.Equal(t, &core.Optional[foo]{Null: true}, optional)
	})

	t.Run("zero type", func(t *testing.T) {
		type foo struct {
			Name string
		}
		zero := new(foo)
		optional := OptionalOrNullPtr(zero)
		assert.Equal(t, &core.Optional[foo]{Value: *zero}, optional)
	})

	t.Run("non-zero type", func(t *testing.T) {
		type foo struct {
			Name string
		}
		value := &foo{Name: "test"}
		optional := OptionalOrNullPtr(value)
		assert.Equal(t, &core.Optional[foo]{Value: *value}, optional)
	})
}
