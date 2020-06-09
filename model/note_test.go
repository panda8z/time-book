/*
go test -v -c -o b.test  github.com/panda8z/time-book/model
go tool test2json -t a.test -test.v -test.run ^TestAddN$
*/
package model

import (
	"fmt"
	"testing"
	"time"

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
