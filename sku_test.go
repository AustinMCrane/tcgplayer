package tcgplayer

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetSKUDetails(t *testing.T) {
	client, err := New(*publicKey, *privateKey)
	require.NoError(t, err)

	skus := []int{1260372, 255924}

	skuDetails, err := client.GetSKUDetails(skus)
	require.NoError(t, err)
	require.Equal(t, len(skuDetails), len(skus))
}
