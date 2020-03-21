package tcgplayer

import (
	"strconv"

	"github.com/pkg/errors"
)

func (client *Client) GetGroups(categoryID int) ([]*Group, error) {
	var groupAPIResponse GroupAPIResponse
	err := get(client, "/catalog/categories/"+strconv.Itoa(categoryID)+"/groups", nil, &groupAPIResponse)
	if err != nil {
		return nil, err
	}

	if len(groupAPIResponse.Results) == 0 {
		return nil, errors.New("did not find any groups")
	}

	return groupAPIResponse.Results, nil
}
