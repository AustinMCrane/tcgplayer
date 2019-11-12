package tcgplayer

import "strconv"

type ProductSKU struct {
	SkuID       int `json:"skuId"`
	ProductID   int `json:"productId"`
	LanguageID  int `json:"languageId"`
	PrintingID  int `json:"printingId"`
	ConditionID int `json:"conditionId"`
}

type productSKUResponse struct {
	response
	Results []*ProductSKU `json:"results"`
}

func (client *Client) ListProductSKUs(productID int) ([]*ProductSKU, error) {
	url := generateURL("/catalog/products/" + strconv.Itoa(productID) + "/skus")

	var productSKUs productSKUResponse
	err := client.get(url, &productSKUs)
	if err != nil {
		return nil, err
	}

	return productSKUs.Results, nil
}

type Field struct {
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
	Value       string `json:"value"`
}

type SKU struct {
	SKUID       int `json:"skuId"`
	ProductID   int `json:"productId"`
	ConditionID int `json:"conditionId"`
}

type ProductDetail struct {
	ProductID      int      `json:"productId"`
	Name           string   `json:"name"`
	CleanName      string   `json:"cleanName"`
	ImageURL       string   `json:"imageUrl"`
	CategoryID     int      `json:"categoryId"`
	GroupID        int      `json:"groupId"`
	URL            string   `json:"url"`
	ModifiedOn     string   `json:"modifiedOn"`
	ExtendedFields []*Field `json:"extendedData"`
	SKUs           []*SKU   `json:"skus"`
}

func (detail *ProductDetail) GetRarity() string {
	rarity := "Not Found"
	for _, field := range detail.ExtendedFields {
		if field.Name == "Rarity" {
			return field.Value
		}
	}

	return rarity
}

// NOTE: change the name to represent the type too ie: Number is weird when
// we are returning a string
func (detail *ProductDetail) GetNumber() string {
	rarity := "Not Found"
	for _, field := range detail.ExtendedFields {
		if field.Name == "Number" {
			return field.Value
		}
	}

	return rarity
}

type productDetailsResponse struct {
	response
	Results []*ProductDetail `json:"results"`
}

func (client *Client) GetProductDetails(productID int, includeSKUs bool, includeExtraFields bool) ([]*ProductDetail, error) {
	url := generateURL("/catalog/products/" + strconv.Itoa(productID))
	if includeSKUs || includeExtraFields {
		url += "?"
		if includeSKUs {
			url += "includeSkus=true"
		}

		if includeSKUs && includeExtraFields {
			url += "&"
		}

		if includeExtraFields {
			url += "getExtendedFields=true"
		}
	}

	var productDetails productDetailsResponse
	err := client.get(url, &productDetails)
	if err != nil {
		return nil, err
	}

	return productDetails.Results, nil
}
