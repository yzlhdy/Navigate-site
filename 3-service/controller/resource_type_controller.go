package controller

import (
	"navigate/service"

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

// func NewResourceTypeController (service service.ResourceTypeService) ResourceTypeController {
// 	return &resourceTypeController{
// 		service: service,
// 	}
// }
