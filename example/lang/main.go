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
	params := tcgplayer.LanguageParams{
		CategoryID: *categoryID,
	}
	client, err := tcgplayer.New(*publicKey, *privateKey)
	if err != nil {
		log.Fatal(err)
	}
	langs, err := client.GetLanguages(&params)
	if err != nil {
		log.Fatal(err)
	}

	for _, l := range langs {
		fmt.Printf("%d,%d,%s,%s\n", *categoryID, l.ID, l.Name, l.Abbreviation)
	}
}
