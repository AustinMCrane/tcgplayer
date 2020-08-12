package tcgplayer

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetCategories(t *testing.T) {
	client, err := New(*publicKey, *privateKey)
	require.NoError(t, err)

	categories, err := client.GetCategories()
	require.NoError(t, err)
	require.NotNil(t, categories)
}
