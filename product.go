package tcgplayer

import (
	"errors"
	"net/url"
	"strconv"
)

const (
	CategoryYugioh = 2
)

type Product struct {
	ID         int    `json:"productId"`
	Name       string `json:"name"`
	CleanName  string `json:"cleanName"`
	ImageURL   string `json:"imageUrl"`
	CategoryID int    `json:"categoryId"`
	GroupID    int    `json:"groupId"`
	URL        string `json:"url"`
	//ModifiedOn time.Time `json:"modifiedOn"`
}

type APIParams interface {
	SetQueryParams(q *url.Values)
}

type ProductParams struct {
	ProductName string `json:"productName"`
	CategoryID  int    `json:"categoryId"`
	GroupName   string `json:"groupName"`
}

type SKU struct {
	SKUID int `json:"skuId"`
}

func (p ProductParams) SetQueryParams(q *url.Values) {
	q.Add("productName", p.ProductName)
	q.Add("categoryId", strconv.Itoa(p.CategoryID))
}

type ProductListAPIResponse struct {
	Response []*Product `json:"results"`
}

func (client *Client) ListAllProducts(params ProductParams) ([]*Product, error) {
	var productAPIResponse ProductListAPIResponse
	u := "/catalog/products?productName=" + url.QueryEscape(params.ProductName) +
		"&categoryId=" + strconv.Itoa(params.CategoryID)
	if params.GroupName != "" {
		u = u + "&groupName=" + url.QueryEscape(params.GroupName)
	}

	err := get(client, u, params, &productAPIResponse)
	if err != nil {
		return nil, err
	}

	return productAPIResponse.Response, nil
}

type SKUListAPIResponse struct {
	Results []*SKU `json:"results"`
}

func (client *Client) ListProductSKUs(productID int) ([]*SKU, error) {
	var skuResponse SKUListAPIResponse
	err := get(client, "/catalog/products/"+strconv.Itoa(productID)+"/skus", nil, &skuResponse)
	if err != nil {
		return nil, err
	}

	return skuResponse.Results, nil
}

type Group struct {
	Name         string `json:"name"`
	Abbreviation string `json:"abbreviation"`
}

type GroupAPIResponse struct {
	Results []*Group `json:"results"`
}

func (client *Client) GetGroupDetails(groupID int) (*Group, error) {
	var groupAPIResponse GroupAPIResponse
	err := get(client, "/catalog/groups/"+strconv.Itoa(groupID), nil, &groupAPIResponse)
	if err != nil {
		return nil, err
	}

	if len(groupAPIResponse.Results) == 0 {
		return nil, errors.New("did not find any groups")
	}

	return groupAPIResponse.Results[0], nil
}
