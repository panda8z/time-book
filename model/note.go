package model

import (
	"log"
	"time"

	"github.com/jinzhu/gorm"
)

// Note is the model for note segment
type Note struct {
	gorm.Model
	State   int32  `json:"state" gorm:"default:1"` // 1 可用 -1 删除 0 停用
	Type    string `json:"type" gorm:"default:'galeone'"`
	Content string `json:"content"`
}

// NoteExistByID : check Note isExist by id
func NoteExistByID(id int) bool {
	var n *Note
	// SELECT id FROM note WHERE id = ? limit 1;
	db.Select("id").Where("id = ?", id).First(&n)
	if n != nil && n.ID > 0 {
		return true
	}
	return false

}

// GetNList is get a slice of Note
func GetNList(pageNum, pageSize int, maps map[string]interface{}) []Note {
	nList := []Note{}
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(nList)
	return nList
}

// GetN is a getter handler
func GetN(id int) (n *Note) {
	// SELECT * FROM users WHERE id = ? LIMIT 1;
	db.Where("id = ?", id).First(&n).Limit(1)
	return n
}

// AddN is Add Note handler
func AddN(nType, content string) (id int, err error) {
	n := &Note{
		Type:    nType,
		Content: content,
	}
	// INSERT INTO note("type", "content") values("note", "2233");
	dbase := db.Create(n)

	db.Model(&n).Related(&n.Model)
	log.Println("after add note id is", n)
	return int(n.ID), dbase.Error // TODO: return with insert ID
}

// DeleteN is delete the note handler
func DeleteN(id int) error {
	// UPDATE users SET state='-1' WHERE id=111 ;
	db.Model(&Note{}).Where("id = ?", id).Update("state", -1).Limit(1)
	return nil
}

// EditN is edit the note handler
func EditN(id int, data map[string]interface{}) error {
	//  UPDATE users SET state='-1'... WHERE id=111 ;
	db.Model(&Note{}).Where("id = ?", id).Updates(data)
	return nil
}

// BeforCreate simple hook for gorm create operation
func (n *Note) BeforCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())
	scope.SetColumn("state", 1)
	return nil
}

// BeforUpdate simple hook for gorm create operation
func (n *Note) BeforUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())
	return nil
}
