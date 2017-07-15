package main

import (
	"log"

	"github.com/funny/fastapi"
	"github.com/privationel/fastapi/sample/module"
)

func main() {
	app := fastapi.New()
	app.Register(1, &module.Service{})
	//
	client, err := app.Dial("tcp", "0.0.0.0:9001")
	if err != nil {
		log.Fatal("setup client failed:", err)
	}

	for i := 0; i < 10; i++ {
		err := client.Send(&module.AddReq{i, i})
		if err != nil {
			log.Fatal("send failed:", err)
		}

		rsp, err := client.Receive()
		if err != nil {
			log.Fatal("recv failed:", err)
		}

		log.Printf("AddRsp: %d", rsp.(*module.AddRsp).C)
	}

}
