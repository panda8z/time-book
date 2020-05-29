package main

import (
	"github.com/panda8z/time-book/pkg/settings"
	"github.com/panda8z/time-book/routers"
)

func main(){

	app := routers.InitRouter()

	app.Listen(settings.HTTPPort)
}