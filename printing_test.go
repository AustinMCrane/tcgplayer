package tcgplayer

import (
	"log"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetPrintings(t *testing.T) {
	client, err := New(*publicKey, *privateKey)
	require.NoError(t, err)

	printings, err := client.GetPrintings(CategoryYugioh)
	require.NoError(t, err)
	for _, p := range printings {
		log.Println(p)
	}
	require.NotNil(t, printings)
}
