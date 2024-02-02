package models

import "time"

// Response data structure
type Response struct {
	StatusCode int                    `json:"status_code"`
	Message    string                 `json:"message"`
	Data       interface{}            `json:"data"`
	Timer      map[string]interface{} `json:"timer"`
}

// Result data structure
type Result struct {
	Code  int
	Data  interface{}
	Error error
	Timer map[string]interface{}
}

// ToResponse is a function to convert Result to Response format
func (r *Result) ToResponse() *Response {
	rsp := Response{StatusCode: r.Code, Message: "Request Success"}

	if r.Data != nil {
		rsp.Data = r.Data
	} else {
		rsp.Data = make(map[string]interface{})
	}

	if r.Timer != nil {
		for idx, timer := range r.Timer {
			r.Timer[idx] = timer.(time.Duration).Milliseconds() // convert to milisecond
		}
		rsp.Timer = r.Timer
	}

	return &rsp
}

func (r *Result) ToResponseError(statusCode int) *Response {
	return &Response{
		StatusCode: statusCode,
		Data:       nil,
		Message:    r.Error.Error(),
	}
}
