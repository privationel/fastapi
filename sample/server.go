package main

import (
	"log"
	"time"

	"github.com/astaxie/beego"
	"github.com/funny/fastapi"
	"github.com/privationel/fastapi/sample/module"
)

func main() {
	// new fastapi
	app := fastapi.New()
	// reguister module
	app.Register(1, &module.Service{})

	server, err := app.Listen("tcp", "0.0.0.0:9001", nil)
	if err != nil {
		log.Fatal("setup server failed:", err)
	}
	go server.Serve()
	beego.Info("server start and listening:0.0.0.0:9001")
	time.Sleep(100 * time.Hour)
}
