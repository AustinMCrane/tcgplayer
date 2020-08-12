package tcgplayer

import (
	"github.com/pkg/errors"
)

type Category struct {
	ID   int    `json:"categoryId"`
	Name string `json:"name"`
}

type CategoryAPIResponse struct {
	Results []*Category `json:"results"`
}

func (client *Client) GetCategories() ([]*Category, error) {
	var resp CategoryAPIResponse
	err := get(client, "/catalog/categories?limit=40", nil, &resp)
	if err != nil {
		return nil, err
	}

	if len(resp.Results) == 0 {
		return nil, errors.New("did not find any categories")
	}

	return resp.Results, nil
}
