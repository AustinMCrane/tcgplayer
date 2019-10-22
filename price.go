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

type ProductPrice struct {
	ProductID      int     `json:"productId"`
	LowPrice       float64 `json:"lowPrice"`
	MidPrice       float64 `json:"midPrice"`
	HighPrice      float64 `json:"highPrice"`
	MarketPrice    float64 `json:"marketPrice"`
	DirectLowPrice float64 `json:"directLowPrice"`
	SubTypeName    string  `json:"subTypeName"`
}

type productPriceResponse struct {
	response
	Results []*ProductPrice `json:"results"`
}

func (response productPriceResponse) GetValidMarketPrices() []*ProductPrice {
	results := []*ProductPrice{}
	for _, r := range response.Results {
		if r.LowPrice > 0.0 && r.MarketPrice > 0.0 {
			results = append(results, r)
		}
	}

	return results
}

func (client *Client) ListProductMarketPrices(productID int) ([]*ProductPrice, error) {
	url := generateURL("pricing/product/" + strconv.Itoa(productID))

	var prices productPriceResponse
	err := client.get(url, &prices)
	if err != nil {
		return nil, err
	}

	return prices.GetValidMarketPrices(), nil
}
