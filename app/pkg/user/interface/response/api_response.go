package response

type ApiResponse struct {
	StatusCode int          `json:"status_code"`
	ErrorCode  *string      `json:"error_code,omitempty"`
	Message    string       `json:"message"`
	Data       *interface{} `json:"data,omitempty"`
	MetaData   *interface{} `json:"metadata,omitempty"`
}

func BuildErrorResponse(responseCode Code) *ApiResponse {
	return &ApiResponse{
		StatusCode: responseCode.HttpStatusCode,
		ErrorCode:  &responseCode.ErrorCode,
		Message:    responseCode.Message,
	}
}

func BuildSuccessResponseWithoutData(responseCode Code) *ApiResponse {
	return &ApiResponse{
		StatusCode: responseCode.HttpStatusCode,
		Message:    responseCode.Message,
	}
}

func BuildSuccessResponseWithData(responseCode Code, data interface{}) *ApiResponse {
	return &ApiResponse{
		StatusCode: responseCode.HttpStatusCode,
		Message:    responseCode.Message,
		Data:       &data,
	}
}

func BuildSuccessResponseWithDataAndMetaData(responseCode Code, data interface{}, metadata interface{}) *ApiResponse {
	return &ApiResponse{
		StatusCode: responseCode.HttpStatusCode,
		Message:    responseCode.Message,
		Data:       &data,
		MetaData:   &metadata,
	}
}
