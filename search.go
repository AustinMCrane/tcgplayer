package tcgplayer

// SearchParams for searching with pagination
type SearchParams struct {
	Sort    string   `json:"sort"`
	Filters []Filter `json:"filters"`

	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

func (params SearchParams) setPagination(page int, limit int) {
	if page >= 0 && limit >= 0 {
		params.Offset = page * limit
		params.Limit = limit
	}
}

// Filter TODO
type Filter struct {
	Name   string   `json:"name"`
	Values []string `json:"values"`
}
