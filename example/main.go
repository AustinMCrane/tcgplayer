package main

import (
	"context"
	"flag"
	"log"

	tcgplayer "github.com/AustinMCrane/tcgplayer"
)

var (
	publicKey  = flag.String("public-key", "", "tcgplayer public key from developer program")
	privateKey = flag.String("private-key", "", "tcgplayer private key from developer program")
	categoryID = flag.Int("category-id", 1, "category id to use")
)

func main() {
	flag.Parse()

	auth, err := tcgplayer.GetAuthTokenProvider(*publicKey, *privateKey)
	if err != nil {
		panic(err)
	}

	// use the auth request editor to add the auth token to the request
	client, err := tcgplayer.NewClientWithResponses(tcgplayer.BaseURL,
		tcgplayer.WithRequestEditorFn(auth))
	if err != nil {
		panic(err)
	}

	// Get Categories
	categories, err := client.GetCategoriesWithResponse(context.Background(), &tcgplayer.GetCategoriesParams{
		// basically all of the categories
		Limit: 100,
	}, []tcgplayer.RequestEditorFn{}...)
	if err != nil {
		panic(err)
	}

	log.Println(categories.JSON200.Results)
}
