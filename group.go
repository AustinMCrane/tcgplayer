package tcgplayer

import (
	"fmt"
	"strconv"

	"github.com/pkg/errors"
)

type GroupParams struct {
	CategoryID int `json:"categoryId"`
	Limit      int `json:"limit"`
	Offset     int `json:"offset"`
}

func (client *Client) GetGroups(params GroupParams) ([]*Group, error) {
	u := "/catalog/categories/" + strconv.Itoa(params.CategoryID) + "/groups"

	if params.Limit != 0 {
		u = u + "?limit=" + fmt.Sprintf("%d", params.Limit)
	}

	if params.Offset != 0 {
		u = u + "&offset=" + fmt.Sprintf("%d", params.Offset)
	}

	var groupAPIResponse GroupAPIResponse
	err := get(client, u, nil, &groupAPIResponse)
	if err != nil {
		return nil, err
	}

	if len(groupAPIResponse.Results) == 0 {
		return nil, errors.New("did not find any groups")
	}

	return groupAPIResponse.Results, nil
}
