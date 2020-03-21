package tcgplayer

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetGroups(t *testing.T) {
	client, err := New(*publicKey, *privateKey)
	require.NoError(t, err)

	groups, err := client.GetGroups(CategoryYugioh)
	require.NoError(t, err)
	require.True(t, len(groups) > 0)
}
