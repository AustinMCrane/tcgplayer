package tcgplayer

import (
	"strconv"

	"github.com/pkg/errors"
)

type Language struct {
	ID           int    `json:"languageId"`
	Name         string `json:"name"`
	Abbreviation string `json:"abbr"`
}
type LanguageParams struct {
	CategoryID int `json:"categoryId"`
}

type LanguageAPIResponse struct {
	Results []*Language `json:"results"`
}

func (client *Client) GetLanguages(params *LanguageParams) ([]*Language, error) {
	u := "/catalog/categories/" + strconv.Itoa(params.CategoryID) + "/languages"

	var langAPIResponse LanguageAPIResponse
	err := get(client, u, nil, &langAPIResponse)
	if err != nil {
		return nil, err
	}

	if len(langAPIResponse.Results) == 0 {
		return nil, errors.New("did not find any languages")
	}

	return langAPIResponse.Results, nil
}
