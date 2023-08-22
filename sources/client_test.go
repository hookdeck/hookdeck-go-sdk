package sources

import (
	context "context"
	"testing"

	hookdeckgosdk "github.com/hookdeck/hookdeck-go-sdk"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClient(t *testing.T) {
	t.Skip("Temporarily included to show an example using optionals")
	client := NewClient()
	source, err := client.UpdateSource(
		context.TODO(),
		"<SOURCE_ID>",
		&hookdeckgosdk.UpdateSourceRequest{
			Name:           hookdeckgosdk.Optional("name"),
			CustomResponse: hookdeckgosdk.Null[hookdeckgosdk.SourceCustomResponse](),
		},
	)
	require.NoError(t, err)
	assert.NotEmpty(t, source)
}
