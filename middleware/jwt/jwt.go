package jwt

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/panda8z/time-book/pkg/e"
	"github.com/panda8z/time-book/utils"
)

// JWT is a middleware 
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}
		code = e.SUCCESS
		token := c.Query("token")
		if token == "" {
			code = e.INVALID_PARAMS
		} else {
			claims, err := utils.ParseToken(token)
			if err != nil {
				code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}
		}
		if code != e.SUCCESS {
			c.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg":  e.Msg(code),
				"data": data,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
