package setting

import (
	"log"
	"time"

	"github.com/go-ini/ini"
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
	HTTPPort int64
	// JwtSecret for server
	JwtSecret string
	// PageSize for application on mobile or web browser
	PageSize int64
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
	HTTPPort = int64(sec.Key("HTTPPort").MustInt(8000))
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustDuration(60)) * time.Second
	WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustDuration(60)) * time.Second

}

func loadApp() {
	sec, err := Cfg.GetSection("app")
	if err != nil {
		log.Fatalf("Fail to get section 'app': %v", err)
	}

	JwtSecret = sec.Key("JWT_SECRET").MustString("!@)*#)!@U#@*!@!)")
	PageSize = int64(sec.Key("PAGE_SIZE").MustInt(10))

}

func loadDatabase() {
	// nothing

}
