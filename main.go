package main

import (
	"log"
	"os"
	"path"

	"github.com/louisevanderlith/droxolite"
	"github.com/louisevanderlith/droxolite/servicetype"
	"github.com/louisevanderlith/vin/routers"

	"github.com/louisevanderlith/vin/core"
)

func main() {
	keyPath := os.Getenv("KEYPATH")
	pubName := os.Getenv("PUBLICKEY")
	//host := os.Getenv("HOST")
	pubPath := path.Join(keyPath, pubName)

	conf, err := droxolite.LoadConfig()

	if err != nil {
		log.Fatal(err)
	}

	// Register with router
	srv := droxolite.NewService(conf.Appname, pubPath, conf.HTTPPort, servicetype.API)

	err = srv.Register()

	if err != nil {
		log.Fatal(err)
	}

	poxy := droxolite.NewEpoxy(srv)
	routers.Setup(poxy)

	core.CreateContext()
	defer core.Shutdown()

	err = poxy.Boot()

	if err != nil {
		log.Fatal(err)
	}
}
