package tcgplayer

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestListProductSKUs(t *testing.T) {
	client, err := New(*publicKey, *privateKey)
	require.NoError(t, err)

	productID := 152944

	skus, err := client.ListProductSKUs(productID)
	require.NoError(t, err)
	require.NotEqual(t, len(skus), 0)
}

func TestProductDetails(t *testing.T) {
	client, err := New(*publicKey, *privateKey)
	require.NoError(t, err)

	search := SearchParams{
		Offset: 0,
		Limit:  10,
		Sort:   "Name ASC",
		Filters: []Filter{
			Filter{Name: "ProductName", Values: []string{"Dark Magician"}},
		},
	}

	productIDs, err := client.SearchCategoryProducts(yugiohCategoryID, search)
	require.NoError(t, err)
	require.NotEqual(t, len(productIDs), 0)

	productID := productIDs[0]

	details, err := client.GetProductDetails(productID, true, true)
	require.NoError(t, err)

	require.Equal(t, details[0].Name, "Dark Magician")
}
