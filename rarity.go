package tcgplayer

import (
	"strconv"

	"github.com/pkg/errors"
)

type Rarity struct {
	ID      int    `json:"rarityId"`
	Name    string `json:"displayText"`
	DBValue string `json:"dbValue"`
}
type RarityParams struct {
	CategoryID int `json:"categoryId"`
}

type RarityAPIResponse struct {
	Results []*Rarity `json:"results"`
}

func (client *Client) GetRarities(params *RarityParams) ([]*Rarity, error) {
	u := "/catalog/categories/" + strconv.Itoa(params.CategoryID) + "/rarities"

	var rarityAPIResponse RarityAPIResponse
	err := get(client, u, nil, &rarityAPIResponse)
	if err != nil {
		return nil, err
	}

	if len(rarityAPIResponse.Results) == 0 {
		return nil, errors.New("did not find any rarities")
	}

	return rarityAPIResponse.Results, nil
}
