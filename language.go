package tcgplayer

import (
	"strconv"

	"github.com/pkg/errors"
)

type Language struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Abbreviation string `json:"abbr"`
}

type LanguageAPIResponse struct {
	Results []*Language `json:"results"`
}

func (client *Client) GetLanguages(categoryID int) ([]*Language, error) {
	var resp LanguageAPIResponse
	err := get(client, "/catalog/categories/"+strconv.Itoa(categoryID)+"/languages", nil, &resp)
	if err != nil {
		return nil, err
	}

	if len(resp.Results) == 0 {
		return nil, errors.New("did not find any languages")
	}

	return resp.Results, nil
}
