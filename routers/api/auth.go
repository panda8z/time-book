package api

import (
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/panda8z/time-book/model"
	"github.com/panda8z/time-book/pkg/e"
	"github.com/panda8z/time-book/utils"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

// Auth is a getter func
func Auth(c *gin.Context) {
	userName := c.Query("username")
	password := c.Query("password")
	valid := validation.Validation{}
	a := auth{Password: password, Username: userName}
	// valid from beego/validation
	ok, _ := valid.Valid(&a)

	data := make(map[string]interface{})
	code := e.INVALID_PARAMS
	if ok {
		isExits := model.CheckAuth(userName, password)
		if isExits {
			token, err := utils.GenerateToken(userName, password)
			if err != nil {
				code = e.ERROR_AUTH_TOKEN
			} else {
				data["token"] = token
				code = e.SUCCESS
			}
		} else {
			code = e.ERROR_AUTH
		}
	}
	c.JSON(http.StatusOK, gin.H{"code": code, "msg": e.Msg(code), "data": data})
}
