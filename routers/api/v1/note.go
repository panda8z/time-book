package v1

import (
	"fmt"
	"log"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/panda8z/time-book/model"
	"github.com/panda8z/time-book/pkg/e"
	"github.com/panda8z/time-book/utils"
)

// Note is a getter handler
// api/v1/note/:id
func Note(c *gin.Context) {
	valid := validation.Validation{}
	var n *model.Note
	var msg string
	id := c.Param("id")
	code := e.INVALID_PARAMS
	valid.Required(id, "id").Message("id 不能为空")
	if valid.HasErrors() {
		code = e.INVALID_PARAMS
		for _, err := range valid.Errors {
			str := fmt.Sprintf("%s", err.Error())
			log.Println(str)
			msg = fmt.Sprintf("%s%s", msg, str)
		}
	} else {
		nID, err := utils.StrTo(id).Int()
		if err != nil {
			code = e.INVALID_PARAMS
		} else {
			if !model.NoteExistByID(nID) {
				code = e.ERROR_NOT_EXIST_ARTICLE
			} else {
				n = model.GetN(nID)
				code = e.SUCCESS
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  fmt.Sprintf("%s %s", e.Msg(code), msg),
		"data": n,
	})
}

// AddN is Add Note handler
func AddN(c *gin.Context) {
	noteType := c.PostForm("type")
	content := c.PostForm("content")

	valid := validation.Validation{}
	valid.Required(noteType, "type").Message("标题不能为空")
	valid.Required(content, "content").Message("内容不能为空")
	code := e.INVALID_PARAMS
	var msg, str string
	if !valid.HasErrors() {
		id, err := model.AddN(noteType, content)
		if err != nil {
			str = fmt.Sprintf("add note faild:%s", err.Error())

		} else {
			str = fmt.Sprintf("note id is %v", id)
		}
		log.Println(str)
		msg = fmt.Sprintf("%s %s", msg, str)
		code = e.SUCCESS

	} else {
		for _, err := range valid.Errors {
			str = fmt.Sprintf("err.key: %s, err.message: %s", err.Key, err.Message)
			log.Println(str)
			msg = fmt.Sprintf("%s %s", msg, str)
		}

	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  fmt.Sprintf("%s %s", msg, e.Msg(code)),
		"data": gin.H{},
	})

}

// DeleteN is delete the note handler
func DeleteN(c *gin.Context) {
	
}

// EditN is edit the note handler
func EditN(c *gin.Context) {

}
