package api

/*
// @Summary Add a new pet to the store
// @Description get string by ID
// @Accept  json
// @Produce  json
// @Param   some_id     path    int     true        "Some ID"
// @Success 200 {string} string    "ok"
// @Failure 400 {object} web.APIError "We need ID!!"
// @Failure 404 {object} web.APIError "Can not find ID"
// @Router /testapi/get-string-by-int/{some_id} [get]
*/

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
// @Summary Auth is a getter func
// @Description get token string by username and password
// @Accept  json
// @Produce  json
// @Param  username  password
// @Success 200
// @Failure 500
// @Router /auth [get]
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

// AddAuth is a register func
// @Summary AddAuth  is a register func
func AddAuth(c *gin.Context) {
	username := c.PostForm("username")

	password := c.PostForm("password")

	valid := validation.Validation{}
	valid.Required(username, "username").Message("用户名不能为空")
	valid.Required(password, "password").Message("密码不能为空")
	code := e.INVALID_PARAMS
	data := make(map[string]interface{})
	if !valid.HasErrors() {
		err := model.AddUser(username, password)
		if err != nil {
			code = e.SUCCESS

		} else {
			code = e.ERROR_AUTH_EXIST
			data["error"] = err.Error()
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.Msg(code),
		"data": data,
	})
}
