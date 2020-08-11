package tcgplayer

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetRarities(t *testing.T) {
	client, err := New(*publicKey, *privateKey)
	require.NoError(t, err)

	rarities, err := client.GetRarities(CategoryYugioh)
	require.NoError(t, err)
	require.NotNil(t, rarities)
}
