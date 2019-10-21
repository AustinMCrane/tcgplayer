package tcgplayer

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	yugiohCategoryID = 2
	notACategoryID   = -1
)

func TestGetCategories(t *testing.T) {
	client, err := New(*publicKey, *privateKey)
	require.NoError(t, err)

	categories, err := client.GetCategories()
	require.NoError(t, err)
	require.NotEqual(t, len(categories), 0)
}

func TestGetCategory(t *testing.T) {
	client, err := New(*publicKey, *privateKey)
	require.NoError(t, err)

	category, err := client.GetCategory(yugiohCategoryID)
	require.NoError(t, err)
	require.Equal(t, category.Name, "YuGiOh")

	// should return not found
	category, err = client.GetCategory(notACategoryID)
	require.Error(t, err)
}

func TestSearchCategoryProducts(t *testing.T) {
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

	skus, err := client.SearchCategoryProducts(yugiohCategoryID, search)
	require.NoError(t, err)
	require.NotEqual(t, len(skus), 0)
}
