package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"

	"github.com/AustinMCrane/tcgplayer"
)

var (
	publicKey   = flag.String("public-key", "", "tcgplayer public key from developer program")
	privateKey  = flag.String("private-key", "", "tcgplayer private key from developer program")
	productName = flag.String("product-name", "Dark Magician", "product name")
	categoryID  = flag.Int("category-id", tcgplayer.CategoryYugioh, "category id to use")
)

type Result struct {
	Name  string                    `json:"name"`
	Set   string                    `json:"set"`
	Price *tcgplayer.SKUMarketPrice `json:"price"`
}

func main() {
	flag.Parse()
	params := tcgplayer.ProductParams{
		ProductName: *productName,
		CategoryID:  *categoryID,
	}
	client, err := tcgplayer.New(*publicKey, *privateKey)
	if err != nil {
		log.Fatalf("error ", err)
	}
	products, err := client.ListAllProducts(params, 0, 100)
	if err != nil {
		log.Fatalf("error ", err)
	}

	if len(products) == 0 {
		log.Fatalf("did not find product with name ", *productName)
	}

	product := products[0]
	skus, err := client.ListProductSKUs(product.ID)
	if err != nil {
		log.Println("error ", err)
	}

	if len(skus) == 0 {
		log.Fatalf("did not find skus with name ", *productName)
	}

	group, err := client.GetGroupDetails(product.GroupID)
	if err != nil {
		log.Fatalf("did not find group ", err)
	}

	prices, err := client.GetSKUPrices([]int{skus[0].SKUID})
	if err != nil {
		log.Println("error ", err)
	}

	if len(prices) == 0 {
		log.Fatalf("did not find prices with name ", *productName)
	}

	r := Result{
		Name:  products[0].Name,
		Set:   group.Name,
		Price: prices[0],
	}
	file, err := json.MarshalIndent(r, "", " ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(file))
}
