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

type Group struct {
	GroupID        int    `json:"groupId"`
	Name           string `json:"name"`
	Abbreviation   string `json:"abbreviation"`
	IsSupplemental bool   `json:"isSupplemental"`
	PublishedOn    string `json:"publishedOn"`
	ModifiedOn     string `json:"modifiedOn"`
	CategoryID     int    `json:"categoryId"`
}

type Condition struct {
	ConditionID  int    `json:"conditionId"`
	Name         string `json:"name"`
	Abbreviation string `json:"abbreviation"`
	DisplayOrder int    `json:"displayOrder"`
}

type ConditionPrice struct {
	Condition Condition   `json:"condition"`
	Price     MarketPrice `json:"price"`
}

type DetailedProductPrice struct {
	// NOTE: is this where rarity should go?
	Rarity string `json:"rarity"`
	Number string `json:"number"`
	// NOTE: i dont think i want this to be added to the requests since i
	// get the card number from the extended details
	Group         *Group            `json:"group"`
	ProductDetail *ProductDetail    `json:"productDetails"`
	ProductPrice  []*ConditionPrice `json:"productPrice"`
}

func (client *Client) GetProductMarketPrice(categoryID int, search SearchParams) ([]*DetailedProductPrice, error) {
	productIDs, err := client.SearchCategoryProducts(categoryID, search)
	if err != nil {
		return nil, err
	}

	details := []*ProductDetail{}
	for _, productID := range productIDs {
		detail, err := client.GetProductDetails(productID, true, true)
		if err != nil {
			return nil, err
		}

		details = append(details, detail...)
	}

	results := []*DetailedProductPrice{}
	for _, detail := range details {
		// NOTE: reference note in struct about Group attribute
		group, err := client.GetGroup(detail.GroupID)
		if err != nil {
			return nil, err
		}

		pricing := []*ConditionPrice{}
		for _, sku := range detail.SKUs {
			if sku.ConditionID == 1 {
				prices, err := client.ListSKUMarketPrices(sku.SKUID)
				if err != nil {
					return nil, err
				}

				if len(prices) > 0 {
					condition, err := client.GetCondition(sku.ConditionID)
					if err != nil {
						return nil, err
					}

					price := ConditionPrice{
						Condition: *condition,
						Price:     *prices[0],
					}

					pricing = append(pricing, &price)
				}

				if len(pricing) > 0 {
					break
				}
			}
		}

		result := DetailedProductPrice{
			Number:        detail.GetNumber(),
			Rarity:        detail.GetRarity(),
			Group:         group,
			ProductPrice:  pricing,
			ProductDetail: detail,
		}

		results = append(results, &result)
	}

	return results, nil
}
