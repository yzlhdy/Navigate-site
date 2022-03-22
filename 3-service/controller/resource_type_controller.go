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

type ResourceTypeController interface {
	ResourceList(c *gin.Context)
	CreateResourceType(c *gin.Context)
	UpdateResourceType(c *gin.Context)
	FindResourceType(c *gin.Context)
	DeleteResourceType(c *gin.Context)
}

type resourceTypeController struct {
	service service.ResourceTypeService
}

func NewResourceTypeController(service service.ResourceTypeService) ResourceTypeController {
	return &resourceTypeController{
		service: service,
	}
}

func (r *resourceTypeController) ResourceList(c *gin.Context) {
	page, limit := utils.GetPageAndLimit(c)
	resourceTypes, total := r.service.ResourceList(page, limit)
	response := helper.BuildResponsePage(200, "success", resourceTypes, total)
	c.JSON(http.StatusOK, response)
}

func (r *resourceTypeController) CreateResourceType(c *gin.Context) {
	var createDto dto.ResourceTypeCreateDto
	c.ShouldBind(&createDto)
	validate, err := helper.BindAndValid(c, &createDto)
	if !validate {
		response := helper.BuildErrorResponse(401, err.Error(), nil, helper.EmptyObjectResponse{})
		c.JSON(http.StatusOK, response)
		return
	}
	resource := r.service.CreateResourceType(createDto)
	response := helper.BuildResponse(200, "success", resource)
	c.JSON(http.StatusOK, response)
}

func (r *resourceTypeController) UpdateResourceType(c *gin.Context) {
	id, error := strconv.Atoi(c.Param("id"))
	if error != nil {
		panic(error)
	}
	var updateDto dto.ResourceTypeUpdateDto
	c.ShouldBind(&updateDto)
	validate, err := helper.BindAndValid(c, &updateDto)
	if !validate {
		response := helper.BuildErrorResponse(401, err.Error(), nil, helper.EmptyObjectResponse{})
		c.JSON(http.StatusOK, response)
		return
	}
	resource := r.service.UpdateResourceType(id, updateDto)
	response := helper.BuildResponse(200, "success", resource)
	c.JSON(http.StatusOK, response)
}

func (r *resourceTypeController) FindResourceType(c *gin.Context) {
	id, error := strconv.Atoi(c.Param("id"))
	if error != nil {
		panic(error)
	}
	resource := r.service.FindResourceType(id)
	response := helper.BuildResponse(200, "success", resource)
	c.JSON(http.StatusOK, response)
}

func (r *resourceTypeController) DeleteResourceType(c *gin.Context) {
	id, error := strconv.Atoi(c.Param("id"))
	if error != nil {
		panic(error)
	}
	resource := r.service.DeleteResourceType(id)
	response := helper.BuildResponse(200, "success", resource)
	c.JSON(http.StatusOK, response)
}
