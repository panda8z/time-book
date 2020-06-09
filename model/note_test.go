package model

import (
	"testing"
	"time"
)

func init() {
	Init()
}
func TestAddN(t *testing.T) {
	nList := []*Note{
		&Note{
			Content: "爱你伊娃年" + string(time.Now().Unix()),
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
