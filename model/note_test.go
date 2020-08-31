/*
go test -v -c -o b.test  github.com/panda8z/time-book/model

./b.test -test.run ^TestDeletN$

go tool test2json -t a.test -test.v -test.run ^TestAddN$
*/
package model

import (
	"fmt"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/panda8z/time-book/pkg/settings"
)

func init() {
	Init()
}

func TestAddN(t *testing.T) {
	nList := []*Note{
		{
			Content: "爱你伊娃年" + fmt.Sprintf("%d", time.Now().Unix()),
			Type:    "note",
		},
	}
	for _, n := range nList {
		id, err := AddN(n.Type, n.Content)
		if err != nil {
			t.Errorf("AddN error is %s", err.Error())
		} else {
			t.Log("Note id is ", id)
		}
	}
}

// ./b.test -test.run ^TestDeletN$
func TestDeletN(t *testing.T) {
	n := &Note{}
	db.Last(n)
	err := DeleteN(int(n.ID))
	if err != nil {
		t.Error(err.Error())
	}
}

// ./b.test -test.run ^TestEditN$
func TestEditN(t *testing.T) {
	n := &Note{}
	db.Last(n)
	EditN(int(n.ID), gin.H{
		"content": "我不是黄铜，我是青铜",
	})
}

func TestGetN(t *testing.T) {
	
}

func TestGetNList(t *testing.T) {

}
