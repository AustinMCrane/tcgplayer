package tcgplayer

import (
	"log"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	darkMagicianSKU = 3546374
)

func TestListSKUMarketPrices(t *testing.T) {
	client, err := New(*publicKey, *privateKey)
	require.NoError(t, err)

	search := SearchParams{
		Offset: 0,
		Limit:  10,
		Sort:   "MinPrice DESC",
		Filters: []Filter{
			Filter{Name: "ProductName", Values: []string{"Dark Magician"}},
		},
	}

	productIDs, err := client.SearchCategoryProducts(yugiohCategoryID, search)
	require.NoError(t, err)

	darkMagicianProductID := productIDs[2]
	skus, err := client.ListProductSKUs(darkMagicianProductID)
	require.NoError(t, err)

	sku := skus[0]

	prices, err := client.ListSKUMarketPrices(sku.SkuID)
	require.NoError(t, err)
	require.NotEqual(t, prices[0].LowPrice, 0)
	log.Println(prices[0].LowPrice)

	// should return not found
	prices, err = client.ListSKUMarketPrices(notACategoryID)
	require.Error(t, err)
}
