package core

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type OptionalRequest struct {
	Id     string            `json:"id"`
	Filter *Optional[string] `json:"filter,omitempty"`
}

func TestOptional(t *testing.T) {
	tests := []struct {
		desc         string
		giveOptional *Optional[any]
		wantBytes    []byte
	}{
		{
			desc: "primitive",
			giveOptional: &Optional[any]{
				Value: "foo",
			},
			wantBytes: []byte(`"foo"`),
		},
		{
			desc: "null primitive",
			giveOptional: &Optional[any]{
				Null: true,
			},
			wantBytes: []byte("null"),
		},
		{
			desc: "object",
			giveOptional: &Optional[any]{
				Value: &OptionalRequest{
					Id: "xyz",
					Filter: &Optional[string]{
						Value: "foo",
					},
				},
			},
			wantBytes: []byte(`{"id":"xyz","filter":"foo"}`),
		},
		{
			desc: "null object",
			giveOptional: &Optional[any]{
				Value: &OptionalRequest{
					Id: "xyz",
					Filter: &Optional[string]{
						Null: true,
					},
				},
			},
			wantBytes: []byte(`{"id":"xyz","filter":null}`),
		},
		{
			desc: "empty object",
			giveOptional: &Optional[any]{
				Value: &OptionalRequest{
					Id: "xyz",
				},
			},
			wantBytes: []byte(`{"id":"xyz"}`),
		},
	}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			bytes, err := json.Marshal(test.giveOptional)
			require.NoError(t, err)
			assert.Equal(t, test.wantBytes, bytes)
		})
	}
}
