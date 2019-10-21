package tcgplayer

import (
	"strconv"
)

type MarketPrice struct {
	SkuID              int     `json:"skuId"`
	LowPrice           float64 `json:"lowPrice"`
	LowestShipping     float64 `json:"lowestShipping"`
	LowestListingPrice float64 `json:"lowestListingPrice"`
	MarketPrice        float64 `json:"marketPrice"`
	DirectLowPrice     float64 `json:"directLowPrice"`
}

type priceResponse struct {
	response
	Results []*MarketPrice `json:"results"`
}

func (client *Client) ListSKUMarketPrices(sku int) ([]*MarketPrice, error) {
	url := generateURL("pricing/sku/" + strconv.Itoa(sku))

	var prices priceResponse
	err := client.get(url, &prices)
	if err != nil {
		return nil, err
	}

	return prices.Results, nil
}
