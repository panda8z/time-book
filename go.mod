module github.com/panda8z/time-book

go 1.14

require (
	github.com/astaxie/beego v1.12.1 // indirect
	github.com/go-sql-driver/mysql v1.5.0 // indirect
	github.com/gofiber/fiber v1.9.6
	github.com/jinzhu/gorm v1.9.12
	github.com/spf13/viper v1.7.0
)

replace (
	github.com/panda8z/time-book/pkg/settings => ./pkg/settings
	github.com/panda8z/time-book/routers => ./routers
)
