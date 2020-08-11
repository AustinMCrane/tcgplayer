package tcgplayer

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetLanguages(t *testing.T) {
	client, err := New(*publicKey, *privateKey)
	require.NoError(t, err)

	languages, err := client.GetLanguages(CategoryYugioh)
	require.NoError(t, err)
	require.NotNil(t, languages)
}
