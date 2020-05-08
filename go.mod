module github.com/panda8z/time-book

go 1.13

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/go-ini/ini v1.55.0 // indirect
	github.com/go-sql-driver/mysql v1.5.0
	github.com/golang/protobuf v1.4.1 // indirect
	github.com/jinzhu/gorm v1.9.12 // indirect
	github.com/jmoiron/sqlx v1.2.0
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	golang.org/x/sys v0.0.0-20200501145240-bc7a7d42d5c3 // indirect
	gopkg.in/go-playground/assert.v1 v1.2.1 // indirect
	gopkg.in/go-playground/validator.v9 v9.29.1 // indirect
)

replace (
	github.com/panda8z/time-book/pkg/setting => ./pkg/setting
	github.com/panda8z/time-book/routers => ./router
)
