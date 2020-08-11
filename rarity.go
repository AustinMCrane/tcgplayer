package tcgplayer

import (
	"strconv"

	"github.com/pkg/errors"
)

type Rarity struct {
	ID          int    `json:"id"`
	DisplayText string `json:"displayText"`
	DBValue     string `json:"dbValue"`
}

type RarityAPIResponse struct {
	Results []*Rarity `json:"results"`
}

func (client *Client) GetRarities(categoryID int) ([]*Rarity, error) {
	var resp RarityAPIResponse
	err := get(client, "/catalog/categories/"+strconv.Itoa(categoryID)+"/rarities", nil, &resp)
	if err != nil {
		return nil, err
	}

	if len(resp.Results) == 0 {
		return nil, errors.New("did not find any rarities")
	}

	return resp.Results, nil
}
