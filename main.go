package main

import (
	//_ "github.com/go-sql-driver/mysql"
	"github.com/panda8z/time-book/routers"
)

//// UserInfo 用户信息
//type UserInfo struct {
//	gorm.Model
//	Name   string
//	Gender string
//	Hobby  string
//}

func main() {
	//db, err := gorm.Open("mysql", "root:123456@(127.0.0.1:3307)/tbook?charset=utf8mb4&parseTime=True&loc=Local")
	//if err != nil {
	//	panic(err)
	//}
	//defer db.Close()
	//
	//// 自动迁移
	//db.AutoMigrate(&UserInfo{})
	//
	//u1 := UserInfo{gorm.Model{ID:3}, "七米", "男", "篮球"}
	//u2 := UserInfo{gorm.Model{ID:4}, "沙河娜扎", "女", "足球"}
	//// 创建记录
	//db.Create(&u1)
	//db.Create(&u2)
	//// 查询
	//var u = new(UserInfo)
	//db.First(u)
	//fmt.Printf("%#v\n", u)
	//
	//var uu UserInfo
	//db.Find(&uu, "hobby=?", "足球")
	//fmt.Printf("%#v\n", uu)
	//
	//// 更新
	//db.Model(&u).Update("hobby", "双色球")
	//// 删除
	//db.Delete(&u)

	r := routers.InitR()
	r.Run(":3000")
	// s := &http.Server{
	// 	Addr:           fmt.Sprintf(":%d", settings.HTTPPort),
	// 	Handler:        r,
	// 	WriteTimeout:   settings.WriteTimeout,
	// 	ReadTimeout:    settings.ReadTimeout,
	// 	MaxHeaderBytes: 1 << 20,
	// }

	// err := s.ListenAndServe()
	// if err != nil {
	// 	log.Printf("Server err: %v", err)
	// }
	// log.Printf("Listening and serving HTTP on %s\n", string(settings.HTTPPort))

}
