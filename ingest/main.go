package main

import (
	"flag"
	"fmt"

	"github.com/AustinMCrane/tcgplayer"
)

var (
	publicKey   = flag.String("public-key", "", "tcgplayer public key from developer program")
	privateKey  = flag.String("private-key", "", "tcgplayer private key from developer program")
	productName = flag.String("product-name", "Dark Magician", "product name")
	categoryID  = flag.Int("category-id", tcgplayer.CategoryYugioh, "category id to use")
)

func main() {
	fmt.Println("vim-go")
}
