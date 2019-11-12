package tcgplayer

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
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

	// should return not found
	prices, err = client.ListSKUMarketPrices(notACategoryID)
	require.Error(t, err)
}

func TestListProductMarketPrices(t *testing.T) {
	client, err := New(*publicKey, *privateKey)
	require.NoError(t, err)

	search := SearchParams{
		Offset: 0,
		Limit:  100,
		Sort:   "MinPrice ASC",
		Filters: []Filter{
			Filter{Name: "ProductName", Values: []string{"Dark Magician Girl"}},
		},
	}

	productIDs, err := client.SearchCategoryProducts(yugiohCategoryID, search)
	require.NoError(t, err)

	for _, p := range productIDs {
		prices, err := client.ListProductMarketPrices(p)
		require.NoError(t, err)
		for _, pr := range prices {
			assert.Greater(t, pr.LowPrice, 0.0)
		}
	}
}

func TestGetProductMarketPrice(t *testing.T) {
	client, err := New(*publicKey, *privateKey)
	require.NoError(t, err)

	search := SearchParams{
		Offset: 0,
		Limit:  100,
		Sort:   "MinPrice ASC",
		Filters: []Filter{
			Filter{Name: "ProductName", Values: []string{"Dark Magician Girl"}},
		},
	}

	prices, err := client.GetProductMarketPrice(yugiohCategoryID, search)
	require.NoError(t, err)
	require.Greater(t, len(prices), 0)

	log.Println(prices)
}
