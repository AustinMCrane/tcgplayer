package tcgplayer

import (
	"fmt"
	"log"
	"strings"

	"github.com/pkg/errors"
)

func (client *Client) GetSKUDetails(skus []int) ([]*SKU, error) {
	var resp SKUListAPIResponse
	skusStr := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(skus)), ","), "[]")
	log.Println(skusStr)
	err := get(client, "/catalog/skus/"+skusStr, nil, &resp)
	if err != nil {
		return nil, err
	}

	if len(resp.Results) == 0 {
		return nil, errors.New("did not find any skus")
	}

	return resp.Results, nil
}
