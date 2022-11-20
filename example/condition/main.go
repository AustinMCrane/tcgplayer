package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/AustinMCrane/tcgplayer"
)

var (
	publicKey  = flag.String("public-key", "", "tcgplayer public key from developer program")
	privateKey = flag.String("private-key", "", "tcgplayer private key from developer program")
	categoryID = flag.Int("category-id", tcgplayer.CategoryYugioh, "category id to use")
)

func main() {
	flag.Parse()
	params := tcgplayer.ConditionParams{
		CategoryID: *categoryID,
	}
	client, err := tcgplayer.New(*publicKey, *privateKey)
	if err != nil {
		log.Fatal(err)
	}
	conds, err := client.GetConditions(&params)
	if err != nil {
		log.Fatal(err)
	}

	for _, c := range conds {
		fmt.Printf("%d,%d,%s,%s,%d\n", *categoryID, c.ID, c.Name, c.Abbreviation, c.DisplayOrder)
	}
}
