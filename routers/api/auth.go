package api

import (
	"github.com/astaxie/beego/validation"
	"github.com/gofiber/fiber"
	models "github.com/panda8z/time-book/model"
	"github.com/panda8z/time-book/pkg/e"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

func Auth(c *fiber.Ctx) {
	userName := c.Query("username")
	password := c.Query("password")

	valid := validation.Validation{}
	a := auth{Password: password, Username: userName}
	ok, _ := valid.Valid(&a)

	data := make(map[string]interface{})
	code := e.INVALID_PARAMS
	if ok {
		isExits := models.CheckAuth(userName, password)
		if isExits {
			token, err := uitls.GenerateToken(userName, password)
			if err != nil {
				code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else {
				data["token"] = token
				code = e.SUCCESS
			}
		} else {
			code = e.ERROR_AUTH
		}
	}
	c.SendStatus(e.SUCCESS)
	c.JSON(map[string]interface{}{"code": code, "msg": e.Msg(code), "data": data})
}
