package tcgplayer

import (
	"fmt"
	"log"
	"strconv"

	"github.com/pkg/errors"
)

type SKUMarketPrice struct {
	SKUID              int     `json:"skuId"`
	LowPrice           float64 `json:"lowPrice"`
	LowestShipping     float64 `json:"lowestShipping"`
	LowestListingPrice float64 `json:"lowestListingPrice"`
	MarketPrice        float64 `json:"marketPrice"`
	DirectLowPrice     float64 `json:"directLowPrice"`
}

type DetailedSKUPrice struct {
	SKU   *SKU            `json:"sku"`
	Price *SKUMarketPrice `json:"price"`
}

func (p *SKUMarketPrice) String() string {
	s := fmt.Sprintf("Low Price: %f\nLowest Shipping: %f\n"+
		"Lowest Listing Price: %f\nMarket Price: %f\nDirect Low Price: %f\n",
		p.LowPrice, p.LowestShipping, p.LowestListingPrice, p.MarketPrice, p.DirectLowPrice)

	return s
}

type SKUMarketPriceListResponse struct {
	Results []*SKUMarketPrice
}

func (client *Client) GetProductPriceWithProductID(productID int) ([]*DetailedSKUPrice, error) {
	skus, err := client.ListProductSKUs(productID)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get skus")
	}

	skuIDs := []int{}
	for _, sku := range skus {
		skuIDs = append(skuIDs, sku.SKUID)
	}

	log.Println("LEN: ", len(skuIDs))
	pr, err := client.GetSKUPrices(skuIDs)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get sku prices")
	}

	skuPrices := []*DetailedSKUPrice{}
	for _, p := range pr {
		for _, sku := range skus {
			if sku.SKUID == p.SKUID {
				skuPrices = append(skuPrices, &DetailedSKUPrice{
					SKU:   sku,
					Price: p,
				})
			}
		}
	}

	return skuPrices, nil
}

func (client *Client) GetProductPrice(categoryID int, cardName string, setName string, rarityName string) (*SKUMarketPrice, error) {
	params := ProductParams{
		GroupName:   setName,
		ProductName: cardName,
		CategoryID:  categoryID,
	}
	products, err := client.ListAllProducts(params)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get prices")
	}

	var product *Product
	for _, p := range products {

		rarity, err := p.GetExtendedData("Rarity")
		if err != nil {
			return nil, errors.Wrap(err, "unable to get product")
		}
		if rarity.Value == rarityName {
			product = p
		}
	}

	if product == nil {
		return nil, errors.New("unable to find product")
	}

	skus, err := client.ListProductSKUs(product.ID)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get skus")
	}

	if len(skus) > 0 {
		pr, err := client.GetSKUPrices([]int{skus[0].SKUID})
		if err != nil {
			return nil, errors.Wrap(err, "unable to get sku prices")
		}

		if len(pr) > 0 {
			return pr[0], nil
		}
	}

	return nil, nil
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
