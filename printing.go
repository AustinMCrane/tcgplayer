package tcgplayer

import (
	"strconv"

	"github.com/pkg/errors"
)

type Printing struct {
	ID           int    `json:"printingId"`
	Name         string `json:"name"`
	DisplayOrder int    `json:"displayOrder"`
}

type PrintingAPIResponse struct {
	Results []*Printing `json:"results"`
}

func (client *Client) GetPrintings(categoryID int) ([]*Printing, error) {
	var resp PrintingAPIResponse
	err := get(client, "/catalog/categories/"+strconv.Itoa(categoryID)+"/printings", nil, &resp)
	if err != nil {
		return nil, err
	}

	if len(resp.Results) == 0 {
		return nil, errors.New("did not find any printings")
	}

	return resp.Results, nil
}
