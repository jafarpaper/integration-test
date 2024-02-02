package request

type Paginate struct {
	Page int `form:"page"`
	Size int `form:"size"`
}

type GetRequest struct {
	Filter    map[string]string `form:"filter"`
	FilterAll string            `form:"filter[]"`
	Sort      map[string]string `form:"sort"`
	Paginate
}
