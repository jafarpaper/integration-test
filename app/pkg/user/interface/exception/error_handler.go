package exception

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"integration-test/app/pkg/user/domain/exception"
	"integration-test/app/pkg/user/interface/response"
	"strings"
)

func RecoverFromPanic(ctx *gin.Context) {
	if err := recover(); err != nil {

		httpResponse := response.BuildErrorResponse(response.GenericServerError)
		ctx.JSON(httpResponse.StatusCode, httpResponse)
	}
}

func HandleError(ctx context.Context, err error) *response.ApiResponse {

	if strings.Contains(err.Error(), exception.NotFoundError.Error()) {
		return response.BuildErrorResponse(response.GenericResourceNotFound)
	}
	if strings.Contains(err.Error(), exception.DuplicateReferenceIdError.Error()) {
		return response.BuildErrorResponse(response.DuplicateError)
	}

	if _, ok := err.(validator.ValidationErrors); ok {
		return response.BuildErrorResponse(response.ValidationError)
	}

	return response.BuildErrorResponse(response.GenericServerError)
}
