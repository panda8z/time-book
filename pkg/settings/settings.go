package settings

import (
	"github.com/spf13/viper"
	"time"
)
const (
	RUN_MODE = "RUN_MODE"
	PAGE_SIZE ="PAGE_SIZE"
	JWT_SECRET = "JWT_SECRET"
	HTTP_PORT = "HTTP_PORT"
	READ_TIMEOUT = "READ_TIMEOUT"
	WRITE_TIMEOUT = "WRITE_TIMEOUT"
	TYPE = "TYPE"
	USER = "USER"
	PASSWORD = "PASSWORD"
	HOST = "HOST"
	NAME = "NAME"
	TABLE_PREFIX = "TABLE_PREFIX"
)
var (
	// Tbvi globle viper
	TbVi *viper.Viper

	RunMode string

	HTTPPort int
	ReadTimeout time.Duration
	WriteTimeout time.Duration

	PageSize int64
	JwtSecret string
)

func init() {
	tbViper := viper.New()
	tbViper.AddConfigPath("./config/")
	tbViper.SetConfigName("server")
	tbViper.ReadInConfig()
	TbVi = tbViper
	loadBase()
	loadServer()
	loadApp()
}


func loadApp() {
	PageSize = TbVi.GetInt64(PAGE_SIZE)
	JwtSecret = TbVi.GetString(JWT_SECRET)
}

func loadServer() {
	HTTPPort = TbVi.GetInt(HTTP_PORT)
	ReadTimeout = TbVi.GetDuration(READ_TIMEOUT)
	WriteTimeout = TbVi.GetDuration(WRITE_TIMEOUT)
}

func loadBase() {
	RunMode = TbVi.GetString(RUN_MODE)
}

