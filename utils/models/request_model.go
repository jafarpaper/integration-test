package models

type PaginationRequest struct {
	First     int    `json:"first"`
	PerPage   int    `json:"rows"`
	SortField string `json:"sort_field"`
	SortOrder int    `json:"sort_order"`
}

type FilterRequest struct {
	PaginationRequest
	Filters map[string]interface{} `json:"filters"`
}

type FilterBody struct {
	Matchmode string `json:"matchmode"`
	Value     string `json:"value"`
}
