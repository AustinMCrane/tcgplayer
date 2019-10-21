package tcgplayer

import (
	"strconv"
)

type Category struct {
	ID                int    `json:"categoryId"`
	Name              string `json:"name"`
	DisplayName       string `json:"displayName"`
	SEOCategoryName   string `json:"seoCategoryName"`
	SealedLabel       string `json:"sealedLabel"`
	NonSealedLabel    string `json:"nonSealedLabel"`
	ConditionGuideURL string `json:"conditionGuideUrl"`
	IsScannable       bool   `json:"isScannable"`
	Popularity        int    `json:"popularity"`

	// NOTE: need to implement custom unmarshal to get the time
	// 	ModifiedOn        time.Time `json:"modifiedOn"`
}

type categoryResponse struct {
	response
	Results []*Category `json:"results"`
}

// GetCategories ...
func (client *Client) GetCategories() ([]*Category, error) {
	url := generateURL("/catalog/categories")

	var categoryResponse categoryResponse
	err := client.get(url, &categoryResponse)
	if err != nil {
		return nil, err
	}

	return categoryResponse.Results, nil
}

// GetCategory ...
func (client *Client) GetCategory(categoryID int) (*Category, error) {
	url := generateURL("/catalog/categories/" + strconv.Itoa(categoryID))

	var categoryResponse categoryResponse
	err := client.get(url, &categoryResponse)
	if err != nil {
		return nil, err
	}

	return categoryResponse.Results[0], nil
}

type productIDResponse struct {
	response
	Results []int `json:"results"`
}

// SearchCategoryProducts for product ids that match the filters
// this also supports paging by setting the limit offset values
func (client *Client) SearchCategoryProducts(categoryID int, search SearchParams) ([]int, error) {
	url := generateURL("catalog/categories/" + strconv.Itoa(categoryID) + "/search")

	var productIDs productIDResponse
	err := client.post(url, search, &productIDs)
	if err != nil {
		return nil, err
	}

	return productIDs.Results, nil
}
