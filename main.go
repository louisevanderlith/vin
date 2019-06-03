package main

import (
	"log"
	"os"
	"path"

	"github.com/louisevanderlith/mango"
	"github.com/louisevanderlith/mango/enums"
	"github.com/louisevanderlith/vin/routers"

	"github.com/astaxie/beego"
	"github.com/louisevanderlith/vin/core"
)

func main() {
	keyPath := os.Getenv("KEYPATH")
	pubName := os.Getenv("PUBLICKEY")
	host := os.Getenv("HOST")
	pubPath := path.Join(keyPath, pubName)

	core.CreateContext()
	defer core.Shutdown()

	// Register with router
	appName := beego.BConfig.AppName
	srv := mango.NewService(appName, pubPath, enums.API)

	port := beego.AppConfig.String("httpport")
	err := srv.Register(port)

	if err != nil {
		log.Print("Register: ", err)
	} else {
		routers.Setup(srv, host)
		beego.Run()
	}
}
