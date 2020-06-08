package utils

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/panda8z/time-book/pkg/settings"
)

// GetPage from the url 
func GetPage(c *gin.Context) int64 {
	page, err := strconv.ParseInt(c.Query("page"), 10, 64)
	if err != nil {
		log.Println(err)
	}
	var result int64
	if page > 0 {
		result = (page - 1) * settings.PageSize
	}

	return result
}
