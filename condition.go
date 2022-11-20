package tcgplayer

import (
	"strconv"

	"github.com/pkg/errors"
)

type Condition struct {
	ID           int    `json:"conditionId"`
	Name         string `json:"name"`
	Abbreviation string `json:"abbreviation"`
	DisplayOrder int    `json:"displayOrder"`
}
type ConditionParams struct {
	CategoryID int `json:"categoryId"`
}

type ConditionAPIResponse struct {
	Results []*Condition `json:"results"`
}

func (client *Client) GetConditions(params *ConditionParams) ([]*Condition, error) {
	u := "/catalog/categories/" + strconv.Itoa(params.CategoryID) + "/conditions"

	var condAPIResponse ConditionAPIResponse
	err := get(client, u, nil, &condAPIResponse)
	if err != nil {
		return nil, err
	}

	if len(condAPIResponse.Results) == 0 {
		return nil, errors.New("did not find any conditions")
	}

	return condAPIResponse.Results, nil
}
