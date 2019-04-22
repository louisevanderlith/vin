package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	"github.com/louisevanderlith/mango"
	"github.com/louisevanderlith/mango/control"
	secure "github.com/louisevanderlith/secure/core"
	"github.com/louisevanderlith/vin/controllers"
)

func Setup(s *mango.Service) {
	ctrlmap := EnableFilter(s)

	lognCtrl := controllers.NewLookupCtrl(ctrlmap)

	beego.Router("/v1/lookup/:vin", lognCtrl, "get:Get")
}

func EnableFilter(s *mango.Service) *control.ControllerMap {
	ctrlmap := control.CreateControlMap(s)

	emptyMap := make(secure.ActionMap)

	ctrlmap.Add("/lookup", emptyMap)

	beego.InsertFilter("/*", beego.BeforeRouter, ctrlmap.FilterAPI)

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:    []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Content-Type"},
		ExposeHeaders:   []string{"Content-Length", "Access-Control-Allow-Origin"},
	}))

	return ctrlmap
}
