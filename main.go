package main

import (
	"github.com/panda8z/time-book/routers"
)

func main() {

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
