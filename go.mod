module github.com/panda8z/time-book

go 1.14

require (
	github.com/astaxie/beego v1.12.1
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-gonic/gin v1.6.3
	github.com/go-playground/validator/v10 v10.3.0 // indirect
	github.com/go-sql-driver/mysql v1.5.0
	github.com/golang/protobuf v1.4.2 // indirect
	github.com/jinzhu/gorm v1.9.12
	github.com/sony/sonyflake v1.0.0 // indirect
	github.com/spf13/viper v1.7.0
	golang.org/x/sys v0.0.0-20200602225109-6fdc65e7d980 // indirect
	google.golang.org/protobuf v1.24.0 // indirect
	gopkg.in/yaml.v2 v2.3.0 // indirect
)

replace (
	github.com/panda8z/time-book/middleware/jwt => ./middleware/jwt
	github.com/panda8z/time-book/model => ./model
	github.com/panda8z/time-book/pkg/e => ./pkg/e
	github.com/panda8z/time-book/pkg/settings => ./pkg/settings
	github.com/panda8z/time-book/routers => ./routers
	github.com/panda8z/time-book/routers/api => ./routers/api
	github.com/panda8z/time-book/utils => ./utils
)
