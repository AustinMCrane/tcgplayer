package tcgplayer

import (
	"errors"
	"fmt"
	"net/url"
	"strconv"
)

const (
	CategoryYugioh = 2
)

type ExtendedData struct {
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
	Value       string `json:"value"`
}

type Product struct {
	ID           int            `json:"productId"`
	Name         string         `json:"name"`
	CleanName    string         `json:"cleanName"`
	ImageURL     string         `json:"imageUrl"`
	CategoryID   int            `json:"categoryId"`
	GroupID      int            `json:"groupId"`
	URL          string         `json:"url"`
	ExtendedData []ExtendedData `json:"extendedData"`
	SKUS         []SKU          `json:"skus"`
	//ModifiedOn time.Time `json:"modifiedOn"`
}

func (p *Product) GetExtendedData(name string) (*ExtendedData, error) {
	for _, e := range p.ExtendedData {
		if e.Name == name {
			return &e, nil
		}
	}

	return nil, errors.New("unable to find extended data for name " + name)
}

type APIParams interface {
	SetQueryParams(q *url.Values)
}

type ProductParams struct {
	ProductName string `json:"productName"`
	CategoryID  int    `json:"categoryId"`
	GroupName   string `json:"groupName"`
	Limit       int    `json:"limit"`
	Offset      int    `json:"offset"`
	Sort        string `json:"sort"`
}

type SKU struct {
	SKUID       int `json:"skuId"`
	ProductID   int `json:"productId"`
	LanguageID  int `json:"languageId"`
	PrintingID  int `json:"printingId"`
	ConditionID int `json:"conditionId"`
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
	u := "/catalog/products?includeSkus=true&getExtendedFields=true&productName=" + url.QueryEscape(params.ProductName) +
		"&categoryId=" + strconv.Itoa(params.CategoryID)
	if params.GroupName != "" {
		u = u + "&groupName=" + url.QueryEscape(params.GroupName)
	}

	if params.Limit != 0 {
		u = u + "&limit=" + fmt.Sprintf("%d", params.Limit)
	}

	if params.Offset != 0 {
		u = u + "&offset=" + fmt.Sprintf("%d", params.Offset)
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
	ID           int    `json:"groupId"`
	CategoryID   int    `json:"categoryId"`
	Name         string `json:"name"`
	Abbreviation string `json:"abbreviation"`
	PublishedOn  string `json:"publishedOn"`
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

func (client *Client) GetProduct(productID int) (*Product, error) {
	var productAPIResponse ProductListAPIResponse
	err := get(client, "/catalog/products/"+strconv.Itoa(productID), nil, &productAPIResponse)
	if err != nil {
		return nil, err
	}

	if len(productAPIResponse.Response) == 0 {
		return nil, errors.New("did not find any groups")
	}

	return productAPIResponse.Response[0], nil
}
