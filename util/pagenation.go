package util

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/panda8z/time-book/pkg/setting"
)

// GetPage : get specify page`s first index number
func GetPage(c *gin.Context) int {
	i := 0
	page, _ := strconv.ParseInt(c.Query("page"))
	if page > 0 {
		i = (page - 1) * setting.PageSize
	}
	return i
}
