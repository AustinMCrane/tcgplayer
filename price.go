package tcgplayer

import (
	"strconv"
)

type SKUMarketPrice struct {
	SKUID              int     `json:"skuId"`
	LowPrice           float64 `json:"lowPrice"`
	LowestShipping     float64 `json:"lowestShipping"`
	LowestListingPrice float64 `json:"lowestListingPrice"`
	MarketPrice        float64 `json:"marketPrice"`
	DirectLowPrice     float64 `json:"directLowPrice"`
}

type SKUMarketPriceListResponse struct {
	Results []*SKUMarketPrice
}

func (client *Client) GetSKUPrices(skus []int) ([]*SKUMarketPrice, error) {
	var priceResponse SKUMarketPriceListResponse
	skuList := ""
	for _, sku := range skus {
		if len(skuList) == 0 {
			skuList = strconv.Itoa(sku)
		} else {
			skuList = skuList + "," + strconv.Itoa(sku)
		}
	}
	err := get(client, "/pricing/sku/"+skuList, nil, &priceResponse)
	if err != nil {
		return nil, err
	}

	return priceResponse.Results, nil
}
