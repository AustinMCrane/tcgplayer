package tcgplayer

import (
	"fmt"
	"log"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetSKUPrices(t *testing.T) {
	client, err := New(*publicKey, *privateKey)
	require.NoError(t, err)

	skus := []int{1260372}
	prices, err := client.GetSKUPrices(skus)
	require.NoError(t, err)
	require.True(t, len(prices) > 0)
}

func TestGetSKUPricesForDarkMagician(t *testing.T) {
	params := ProductParams{
		ProductName: "Dark Magician",
		CategoryID:  CategoryYugioh,
	}
	client, err := New(*publicKey, *privateKey)
	require.NoError(t, err)

	products, err := client.ListAllProducts(params)
	require.NoError(t, err)
	require.True(t, len(products) > 0)
	product := products[0]
	log.Println(fmt.Sprintf("%+v", product))

	skus, err := client.ListProductSKUs(products[0].ID)
	require.NoError(t, err)
	require.True(t, len(skus) > 0)

	prices, err := client.GetSKUPrices([]int{skus[0].SKUID})
	require.NoError(t, err)
	require.True(t, len(prices) > 0)
}

func TestGetProductPriceWithProductID(t *testing.T) {
	productID := 21876
	client, err := New(*publicKey, *privateKey)
	require.NoError(t, err)

	prices, err := client.GetProductPriceWithProductID(productID)
	require.NoError(t, err)
	for _, p := range prices {
		log.Println(fmt.Sprintf("%+v %+v", p.SKU, p.Price))
	}
}
