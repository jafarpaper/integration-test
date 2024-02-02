package response

type MetaData struct {
	ClientProperties map[string]interface{} `json:"client_properties,omitempty"`
	Pagination       Pagination             `json:"pagination"`
}
type Pagination struct {
	TotalPages int `json:"total_pages"`
	Page       int `json:"page"`
	Size       int `json:"size"`
	TotalData  int `json:"total_data"`
}
