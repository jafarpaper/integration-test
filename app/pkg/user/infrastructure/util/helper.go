package util

import "integration-test/app/pkg/user/domain/model/response"

const (
	UnformattedValueRenderOption = "UNFORMATTED_VALUE"
	UserEnteredValueInputOption  = "USER_ENTERED"
)

func Transpose(slice [][]interface{}) [][]interface{} {
	xl := len(slice[0])
	yl := len(slice)
	result := make([][]interface{}, xl)
	for i := range result {
		result[i] = make([]interface{}, yl)
	}
	for i := 0; i < xl; i++ {
		for j := 0; j < yl; j++ {
			result[i][j] = slice[j][i]
		}
	}
	return result
}

func SetPagination(page int, size int, total int) (paging *response.Pagination) {
	var totalPages = 1
	if total > 0 && total > size {
		mod := total % size
		totalPages = (total - mod) / size
		if mod > 0 {
			totalPages = totalPages + 1
		}
	}

	return &response.Pagination{
		TotalPages: totalPages,
		Page:       page,
		Size:       size,
		TotalData:  total,
	}
}

func GetMetaData(page int, size int, total int) (paging *response.MetaData) {
	data := SetPagination(page, size, total)

	return &response.MetaData{Pagination: *data}
}
