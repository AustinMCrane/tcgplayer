package tcgplayer

import "strconv"

type groupResponse struct {
	response
	Results []*Group `json:"results"`
}

func (client *Client) GetGroup(groupID int) (*Group, error) {
	url := generateURL("/catalog/groups/" + strconv.Itoa(groupID))

	var groupResponse groupResponse
	err := client.get(url, &groupResponse)
	if err != nil {
		return nil, err
	}

	return groupResponse.Results[0], nil
}
