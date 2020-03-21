package tcgplayer

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestListAllPRoducts(t *testing.T) {
	params := ProductParams{
		ProductName: "Dark Magician",
		CategoryID:  CategoryYugioh,
	}
	client, err := New(*publicKey, *privateKey)
	require.NoError(t, err)

	products, err := client.ListAllProducts(params)
	require.NoError(t, err)
	require.True(t, len(products) > 0)
}

func TestListProductSKUs(t *testing.T) {
	client, err := New(*publicKey, *privateKey)
	require.NoError(t, err)

	productID := 152944
	skus, err := client.ListProductSKUs(productID)
	require.NoError(t, err)
	require.True(t, len(skus) > 0)
}

func TestGetGroupDetails(t *testing.T) {
	client, err := New(*publicKey, *privateKey)
	require.NoError(t, err)

	groupID := 298
	group, err := client.GetGroupDetails(groupID)
	require.NoError(t, err)
	require.NotNil(t, group)
}
