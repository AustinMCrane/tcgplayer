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

type ProductDetail struct {
	ProductID  int    `json:"productId"`
	Name       string `json:"name"`
	CleanName  string `json:"cleanName"`
	ImageURL   string `json:"imageUrl"`
	CategoryID int    `json:"categoryId"`
	GroupID    int    `json:"groupId"`
	URL        string `json:"url"`
	ModifiedOn string `json:"modifiedOn"`
}

type productDetailsResponse struct {
	response
	Results []*ProductDetail `json:"results"`
}

func (client *Client) GetProductDetails(productID int) ([]*ProductDetail, error) {
	url := generateURL("/catalog/products/" + strconv.Itoa(productID))

	var productDetails productDetailsResponse
	err := client.get(url, &productDetails)
	if err != nil {
		return nil, err
	}

	return productDetails.Results, nil
}
