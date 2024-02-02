package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"integration-test/app/pkg/user/domain/model/request"
	"integration-test/app/pkg/user/domain/usecase"
	"integration-test/app/pkg/user/infrastructure/util"
	"integration-test/app/pkg/user/interface/exception"
	"integration-test/app/pkg/user/interface/response"
	"strconv"
)

type UserController struct {
	userUseCase      usecase.UserUseCase
	requestValidator *validator.Validate
}

func NewHttpUserController(useCase usecase.UserUseCase, requestValidator *validator.Validate) *UserController {
	return &UserController{
		userUseCase:      useCase,
		requestValidator: requestValidator,
	}
}

func (e *UserController) GetUser(ctx *gin.Context) {
	var (
		httpResponse *response.ApiResponse = nil
		err          error                 = nil
	)
	defer exception.RecoverFromPanic(ctx)

	req := payloadData(ctx, err)

	err = e.requestValidator.Struct(req)
	if err != nil {
		httpResponse = exception.HandleError(ctx, err)
		ctx.JSON(httpResponse.StatusCode, httpResponse)
		return
	}

	data, total, err := e.userUseCase.Get(ctx, req)
	if err != nil {
		httpResponse = exception.HandleError(ctx, err)
		ctx.JSON(httpResponse.StatusCode, httpResponse)
		return
	}

	metadata := util.GetMetaData(req.Page, req.Size, total)

	httpResponse = response.BuildSuccessResponseWithDataAndMetaData(response.Ok, data, metadata)
	ctx.JSON(httpResponse.StatusCode, httpResponse)
}

func (e *UserController) FindUserById(ctx *gin.Context) {
	var (
		httpResponse *response.ApiResponse = nil
		err          error                 = nil
		id                                 = ctx.Param("id")
	)
	defer exception.RecoverFromPanic(ctx)

	parseId, err := uuid.Parse(id)
	if err != nil {
		httpResponse = exception.HandleError(ctx, err)
		ctx.JSON(httpResponse.StatusCode, httpResponse)
		return
	}

	data, err := e.userUseCase.FindById(ctx, parseId)
	if err != nil {
		httpResponse = exception.HandleError(ctx, err)
		ctx.JSON(httpResponse.StatusCode, httpResponse)
		return
	}

	httpResponse = response.BuildSuccessResponseWithData(response.Ok, data)
	ctx.JSON(httpResponse.StatusCode, httpResponse)
}

func (e *UserController) Create(ctx *gin.Context) {
	var (
		httpResponse *response.ApiResponse = nil
		err          error                 = nil
	)
	defer exception.RecoverFromPanic(ctx)

	var req request.UserRequest
	ctx.BindJSON(&req)

	err = e.requestValidator.Struct(req)
	if err != nil {
		httpResponse = exception.HandleError(ctx, err)
		ctx.JSON(httpResponse.StatusCode, httpResponse)
		return
	}

	err = e.userUseCase.Create(ctx, req)
	if err != nil {
		httpResponse = exception.HandleError(ctx, err)
		ctx.JSON(httpResponse.StatusCode, httpResponse)
		return
	}

	httpResponse = response.BuildSuccessResponseWithoutData(response.Created)
	ctx.JSON(httpResponse.StatusCode, httpResponse)
}

func (e *UserController) UpdateUser(ctx *gin.Context) {
	var (
		httpResponse *response.ApiResponse = nil
		err          error                 = nil
		id                                 = ctx.Param("id")
	)
	defer exception.RecoverFromPanic(ctx)

	var req request.UserRequest
	ctx.BindJSON(&req)

	err = e.requestValidator.Struct(req)
	if err != nil {
		httpResponse = exception.HandleError(ctx, err)
		ctx.JSON(httpResponse.StatusCode, httpResponse)
		return
	}

	parseId, err := uuid.Parse(id)
	if err != nil {
		httpResponse = exception.HandleError(ctx, err)
		ctx.JSON(httpResponse.StatusCode, httpResponse)
		return
	}

	err = e.userUseCase.Update(ctx, req, parseId)
	if err != nil {
		httpResponse = exception.HandleError(ctx, err)
		ctx.JSON(httpResponse.StatusCode, httpResponse)
		return
	}

	httpResponse = response.BuildSuccessResponseWithoutData(response.Ok)
	ctx.JSON(httpResponse.StatusCode, httpResponse)
}

func (e *UserController) DeleteUser(ctx *gin.Context) {
	var (
		httpResponse *response.ApiResponse = nil
		err          error                 = nil
		id                                 = ctx.Param("id")
	)
	defer exception.RecoverFromPanic(ctx)

	parseId, err := uuid.Parse(id)
	if err != nil {
		httpResponse = exception.HandleError(ctx, err)
		ctx.JSON(httpResponse.StatusCode, httpResponse)
		return
	}

	err = e.userUseCase.Delete(ctx, parseId)
	if err != nil {
		httpResponse = exception.HandleError(ctx, err)
		ctx.JSON(httpResponse.StatusCode, httpResponse)
		return
	}

	httpResponse = response.BuildSuccessResponseWithoutData(response.Ok)
	ctx.JSON(httpResponse.StatusCode, httpResponse)
}

func payloadData(context *gin.Context, err error) request.GetRequest {
	var payload request.GetRequest
	payload.Filter = context.QueryMap("filter")
	payload.Sort = context.QueryMap("sort")
	payload.FilterAll = context.Query("filter[]")

	pageParam := context.Query("page")
	page, err := strconv.Atoi(pageParam)
	if err != nil {
		page = 1
	}
	payload.Page = page

	sizeParam := context.Query("size")
	size, err := strconv.Atoi(sizeParam)
	if err != nil {
		size = 100
	}
	payload.Size = size
	return payload
}
