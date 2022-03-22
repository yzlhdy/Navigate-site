package controller

import (
	"navigate/dto"
	"navigate/helper"
	"navigate/service"
	"navigate/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ResourceController interface {
	// 增删改查
	GetAll(ctx *gin.Context)
	Get(ctx *gin.Context)
	Create(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Update(ctx *gin.Context)
}

type resourceController struct {
	service service.ResourceService
}

func NewResourceController(service service.ResourceService) ResourceController {
	return &resourceController{
		service: service,
	}
}

func (s *resourceController) Create(ctx *gin.Context) {
	var createDto dto.ResourceCreateDto
	ctx.ShouldBind(&createDto)
	validate, err := helper.BindAndValid(ctx, &createDto)
	if !validate {
		response := helper.BuildErrorResponse(401, err.Error(), nil, helper.EmptyObjectResponse{})
		ctx.JSON(http.StatusOK, response)
		return
	}
	result := s.service.Create(createDto)
	response := helper.BuildResponse(200, "success", result)
	ctx.JSON(http.StatusOK, response)
}

func (s *resourceController) Update(ctx *gin.Context) {
	var updateDto dto.ResourceUpdateDto
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		response := helper.BuildErrorResponse(401, err.Error(), nil, helper.EmptyObjectResponse{})
		ctx.JSON(http.StatusOK, response)
		return
	}
	ctx.ShouldBind(&updateDto)
	validate, err := helper.BindAndValid(ctx, &updateDto)
	if !validate {
		response := helper.BuildErrorResponse(401, err.Error(), nil, helper.EmptyObjectResponse{})
		ctx.JSON(http.StatusOK, response)
		return
	}
	result := s.service.Update(id, updateDto)
	response := helper.BuildResponse(200, "success", result)
	ctx.JSON(http.StatusOK, response)
}

func (s *resourceController) Get(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		response := helper.BuildErrorResponse(401, err.Error(), nil, helper.EmptyObjectResponse{})
		ctx.JSON(http.StatusOK, response)
		return
	}
	result := s.service.Get(id)
	response := helper.BuildResponse(200, "success", result)
	ctx.JSON(http.StatusOK, response)
}

func (s *resourceController) GetAll(ctx *gin.Context) {
	page, limit := utils.GetPageAndLimit(ctx)
	result, total := s.service.GetAll(page, limit)
	response := helper.BuildResponsePage(200, "success", result, total)
	ctx.JSON(http.StatusOK, response)
}

func (s *resourceController) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		response := helper.BuildErrorResponse(401, err.Error(), nil, helper.EmptyObjectResponse{})
		ctx.JSON(http.StatusOK, response)
		return
	}
	result := s.service.Delete(id)
	response := helper.BuildResponse(200, "success", result)
	ctx.JSON(http.StatusOK, response)
}
