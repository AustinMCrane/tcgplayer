package tcgplayer

import "strconv"

type conditionResponse struct {
	response
	Results []*Condition `json:"results"`
}

func (client *Client) GetCondition(conditionID int) (*Condition, error) {
	url := generateURL("catalog/categories/" + strconv.Itoa(conditionID) + "/conditions")

	var conditionResponse conditionResponse
	err := client.get(url, &conditionResponse)
	if err != nil {
		return nil, err
	}

	return conditionResponse.Results[0], nil
}
