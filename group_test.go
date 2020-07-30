package tcgplayer

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetGroups(t *testing.T) {

	client, err := New(*publicKey, *privateKey)
	require.NoError(t, err)

	params := GroupParams{
		CategoryID: CategoryYugioh,
		Limit:      10,
		Offset:     10,
	}

	groups, err := client.GetGroups(params)
	require.NoError(t, err)
	require.True(t, len(groups) > 0)
}
