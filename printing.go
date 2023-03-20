package tcgplayer

import (
	"strconv"

	"github.com/pkg/errors"
)

type PrintingParams struct {
	CategoryID int `json:"categoryId"`
}

type Printing struct {
	ID   int    `json:"printingId"`
	Name string `json:"name"`
}

type PrintingAPIResponse struct {
	Results []*Printing `json:"results"`
}

func (client *Client) GetPrinting(params PrintingParams) ([]*Printing, error) {
	u := "/catalog/categories/" + strconv.Itoa(params.CategoryID) + "/printings"

	var printingAPIResponse PrintingAPIResponse
	err := get(client, u, nil, &printingAPIResponse)
	if err != nil {
		return nil, err
	}

	if len(printingAPIResponse.Results) == 0 {
		return nil, errors.New("did not find any printings")
	}

	return printingAPIResponse.Results, nil
}
