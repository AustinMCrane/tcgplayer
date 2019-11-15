package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/AustinMCrane/tcgplayer"
)

var productName = flag.String("n", "Dark Magician", "name of the product to search for")
var categoryID = flag.Int("c", 2, "category of the product to lookup (1) Magic (2) Yugioh")
var publicKey = flag.String("p", "", "tcgplayer public key")
var privateKey = flag.String("s", "", "tcgplayer private key")

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func printPricing(p *tcgplayer.MarketPrice) {
	if p != nil {
		str := fmt.Sprintf("\nLow Price: $%0.2f\nLowest Shipping: $%0.2f\nLowest Listing Price: $%0.2f\nMarket Price: $%0.2f\nDirect Low Price: $%0.2f",
			p.LowPrice, p.LowestShipping, p.LowestListingPrice, p.MarketPrice, p.DirectLowPrice)
		log.Println(str)
	}
}

func printProductDetails(p *tcgplayer.ProductDetail) {
	if p != nil {
		str := fmt.Sprintf("Name: %s\nClean Name: %s\nImage URL: %s\n", p.Name, p.CleanName, p.ImageURL)
		log.Println(str)
	}
}

func main() {
	flag.Parse()

	client, err := tcgplayer.New(*publicKey, *privateKey)
	handleError(err)

	search := tcgplayer.SearchParams{
		Offset: 0,
		Limit:  10,
		Sort:   "MinPrice DESC",
		Filters: []tcgplayer.Filter{
			tcgplayer.Filter{Name: "ProductName", Values: []string{*productName}},
		},
	}

	productIDs, err := client.SearchCategoryProducts(*categoryID, search)
	handleError(err)

	for _, p := range productIDs {
		details, err := client.GetProductDetails(p)
		handleError(err)
		printProductDetails(details[0])
	}

	productID := productIDs[0]

	skus, err := client.ListProductSKUs(productID)

	sku := skus[0]
	prices, err := client.ListSKUMarketPrices(sku.SkuID)
	handleError(err)

	price := prices[0]
	printPricing(price)
}
