package utils

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetPageAndLimit(c *gin.Context) (int, int) {
	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "10")
	pageData, err := strconv.Atoi(page)
	limitData, error := strconv.Atoi(limit)
	if err != nil || error != nil {
		pageData = 1
		limitData = 10
	}
	return pageData, limitData
}
