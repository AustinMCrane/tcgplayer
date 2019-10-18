package tcgplayer

import (
	"flag"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	publicKey  = flag.String("public-key", "", "tcgplayer public key from developer program")
	privateKey = flag.String("private-key", "", "tcgplayer private key from developer program")
)

func TestMain(m *testing.M) {
	flag.Parse()
	os.Exit(m.Run())
}

func TestNew(t *testing.T) {
	_, err := New(*publicKey, *privateKey)
	require.NoError(t, err)
}
