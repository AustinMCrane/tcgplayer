package tcgplayer

import (
	"fmt"
	"log"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestListAllPRoducts(t *testing.T) {
	params := ProductParams{
		GroupName:   "Duel Overload",
		ProductName: "Crystron Halqifibrax",
		CategoryID:  CategoryYugioh,
	}
	client, err := New(*publicKey, *privateKey)
	require.NoError(t, err)

	products, err := client.ListAllProducts(params)
	require.NoError(t, err)
	require.True(t, len(products) == 1)

	product := products[0]
	require.Equal(t, product.Name, params.ProductName)
}

func TestListAllPRoductsPage(t *testing.T) {
	params := ProductParams{
		CategoryID: CategoryYugioh,
		Limit:      10,
		Offset:     10,
	}
	client, err := New(*publicKey, *privateKey)
	require.NoError(t, err)

	products, err := client.ListAllProducts(params)
	require.NoError(t, err)
	require.True(t, len(products) == params.Limit)
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

func TestPriceOfNeedleFiber(t *testing.T) {
	params := ProductParams{
		GroupName:   "Duel Overload",
		ProductName: "Crystron Halqifibrax",
		CategoryID:  CategoryYugioh,
	}
	client, err := New(*publicKey, *privateKey)
	require.NoError(t, err)

	products, err := client.ListAllProducts(params)
	require.NoError(t, err)
	require.True(t, len(products) == 1)

	product := products[0]
	require.Equal(t, product.Name, params.ProductName)

	skus, err := client.ListProductSKUs(product.ID)
	require.NoError(t, err)
	require.True(t, len(skus) > 0)

	prices, err := client.GetSKUPrices([]int{skus[0].SKUID})
	require.NoError(t, err)
	require.True(t, len(prices) > 0)
}

func TestGetConditions(t *testing.T) {
	client, err := New(*publicKey, *privateKey)
	require.NoError(t, err)

	conditions, err := client.GetConditions(CategoryYugioh)
	require.NoError(t, err)

	for _, c := range conditions {
		log.Println(fmt.Sprintf("%+v", c))
	}
}
