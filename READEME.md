# time-book 时光笔记本

## 需求分析

## 项目分析

## 代码组织和系统架构

### 系统架构

### 代码组织

## 统一配置组件

### add `ini`

`go get -u github.com/go-ini/ini`

### 配置文件解析包： `setting`

```go
package settings

import (
	"log"
	"time"

	"gopkg.in/ini.v1"
)

// Server setting model
type Server struct {
	RunMode      string
	HTTPPort     int // HttpPort custom port
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

// Database setting model
type Database struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
}

var (
	// DatabaseSetting global object
	DatabaseSetting = &Database{}
	// ReadTimeout count
	ReadTimeout time.Duration
	// WriteTimeout count
	WriteTimeout time.Duration
	// Cfg is global ini.File obj
	Cfg *ini.File
	// RunMode server mod
	RunMode string
	// HTTPPort of server
	HTTPPort int
	// JwtSecret for server
	JwtSecret string
	// PageSize for application on mobile or web browser
	PageSize int
)

// Setup initialize the configuration instance
func init() {
	var err error
	Cfg, err = ini.Load("config/config.ini")
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'config/config.ini': %v", err)
	}

	loadBase()
	loadServer()
	loadApp()
	loadDatabase()
}

func loadBase() {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}

func loadServer() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("Fail to get section 'server': %v", err)
	}
	HTTPPort = sec.Key("HTTPPort").MustInt(8000)
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustDuration(60)) * time.Second
	WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustDuration(60)) * time.Second

}

func loadApp() {
	sec, err := Cfg.GetSection("app")
	if err != nil {
		log.Fatalf("Fail to get section 'app': %v", err)
	}

	JwtSecret = sec.Key("JWT_SECRET").MustString("!@)*#)!@U#@*!@!)")
	PageSize = sec.Key("PAGE_SIZE").MustInt(10)
}

func loadDatabase() {
    // nothing
}


```

## 错误码和错误信息定义

### 错误码定义

### 错误信息定义

## 基础路由测试

### add `gin`

` go get -u github.com/gin-gonic/gin`

### 路由包定义

### 服务器基础设置

```bash
.
├── LICENSE
├── READEME.md
├── config
│   └── config.ini
├── docs
├── go.mod
├── go.sum
├── main.go
├── middleware
├── model
│   ├── note.go
│   ├── notebook.sql
│   ├── user.go
│   └── userinfo.go
├── pkg
│   ├── errors
│   │   ├── code.go
│   │   └── msg.go
│   └── setting
│       └── setting.go
├── routers
│   └── router.go
├── runtime
├── util
│   └── pagenation.go
└── view

12 directories, 20 files

```

## 数据库工具链

### add `grom` `go-sql-driver/mysql`

`go get -u github.com/jinzhu/gorm`

` go get -u github.com/go-sql-driver/mysql`

### note 定义
### user 定义
### auth 定义

## 
