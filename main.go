package main

import (
	"fmt"
	"net/http"

	router "github.com/panda8z/time-book/routers"
)

func main() {
	r := router.InitRouter()
	s := &http.Server{
		Addr:           fmt.Sprint(":%d", setting.HTTPPort),
		Handler:        r,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()

}
