package tcgplayer

import (
	"github.com/pkg/errors"
)

type Category struct {
	ID         int    `json:"categoryId"`
	Name       string `json:"displayName"`
	ModifiedOn string `json:"modifiedOn"`
}

type CategoryAPIResponse struct {
	Results []*Category `json:"results"`
}

func (client *Client) GetCategories() ([]*Category, error) {
	u := "/catalog/categories?limit=100"

	var categoryAPIResponse CategoryAPIResponse
	err := get(client, u, nil, &categoryAPIResponse)
	if err != nil {
		return nil, err
	}

	if len(categoryAPIResponse.Results) == 0 {
		return nil, errors.New("did not find any categories")
	}

	return categoryAPIResponse.Results, nil
}
