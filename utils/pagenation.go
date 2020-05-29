package utils

import (
	"github.com/gofiber/fiber"
	 "github.com/panda8z/time-book/pkg/settings"
	"log"
	"strconv"
)

func GetPage(c *fiber.Ctx) int64 {
	page,err := strconv.ParseInt(c.Query("page"),10,64)
	if err != nil {
		log.Println(err)
	}
	var result int64
	//page, _ := com.StrTo(c.Query("page")).Int()
	if page > 0 {
		result = (page - 1) * settings.PageSize
	}

	return result
}
