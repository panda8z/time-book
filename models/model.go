package models

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	"github.com/panda8z/time-book/pkg/setting"
)

var db *gorm.DB

// Model the common word
type Model struct {
	ID         int `gorm:"primary_key" json:id`
	CreatedOn  int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
}

func init() {
	var (
		err                                               error
		dbType, dbName, user, password, host, tablePrefix string
	)

	sec, err := setting.Cfg.GetSection("database")
	if err != nil {
		log.Fatal(2, "Fail to get section 'server': %v", err)
	}
	user = sec.Key("TYPE").MustString("mysql")
	dbType = sec.Key("USER").MustString("root")
	password = sec.Key("PASSWORD").MustString("123456")
	host = sec.Key("HOST").MustString("127.0.0.1")
	dbName = sec.Key("NAME").MustString("timebook")
	tablePrefix = sec.Key("TABLE_PREFIX").MustString("tb_")

	db, err = gorm.Open(dbType,
		fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
			user,
			password,
			host,
			dbName))

	if err != nil {
		log.Println(err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}

	db.SingularTable(true)
	db.LogMode(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}

// CloseDB in the end
func CloseDB() {
	defer db.Close()
}
