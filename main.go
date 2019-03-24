package main

import (
	"log"
	"os"

	"github.com/louisevanderlith/mango"
	"github.com/louisevanderlith/mango/enums"
	"github.com/louisevanderlith/vin/routers"

	"github.com/astaxie/beego"
	"github.com/louisevanderlith/vin/core"
)

func main() {
	mode := os.Getenv("RUNMODE")

	core.CreateContext()
	defer core.Shutdown()

	// Register with router
	appName := beego.BConfig.AppName
	srv := mango.NewService(mode, appName, enums.API)

	port := beego.AppConfig.String("httpport")
	err := srv.Register(port)

	if err != nil {
		log.Print("Register: ", err)
	} else {
		routers.Setup(srv)
		beego.Run()
	}
}
