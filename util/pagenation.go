package util

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/panda8z/time-book/pkg/setting"
)

// Page : get specify page`s first index number
func Page(c *gin.Context) int64 {
	var i int64
	page, _ := strconv.ParseInt(c.Query("page"),10,64)
	if page > 0 {
		i = (page - 1) * setting.PageSize
	}
	return i
}
