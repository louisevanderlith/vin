package routers

import (
	"fmt"
	"strings"

	"github.com/louisevanderlith/secure/core/roletype"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	"github.com/louisevanderlith/mango"
	"github.com/louisevanderlith/mango/control"
	secure "github.com/louisevanderlith/secure/core"
	"github.com/louisevanderlith/vin/controllers"
)

func Setup(s *mango.Service, host string) {
	ctrlmap := EnableFilter(s, host)

	adminCtrl := controllers.NewAdminCtrl(ctrlmap)
	beego.Router("/v1/admin/:key", adminCtrl, "get:GetByKey")
	beego.Router("/v1/admin/all/:pagesize", adminCtrl, "get:Get")
	beego.Router("/v1/lookup/:vin", controllers.NewLookupCtrl(ctrlmap), "get:Get")
	beego.Router("/v1/validate/:vin", controllers.NewValidateCtrl(ctrlmap), "get:Get")
}

func EnableFilter(s *mango.Service, host string) *control.ControllerMap {
	ctrlmap := control.CreateControlMap(s)

	emptyMap := make(secure.ActionMap)
	emptyMap["GET"] = roletype.User
	emptyMap["POST"] = roletype.Owner

	ctrlmap.Add("/v1/lookup", emptyMap)
	ctrlmap.Add("/v1/validate", emptyMap)

	adminMap := make(secure.ActionMap)
	adminMap["GET"] = roletype.Admin
	adminMap["POST"] = roletype.Admin
	ctrlmap.Add("/v1/admin", adminMap)

	beego.InsertFilter("/v1/*", beego.BeforeRouter, ctrlmap.FilterAPI, false)
	allowed := fmt.Sprintf("https://*%s", strings.TrimSuffix(host, "/"))

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins: []string{allowed},
		AllowMethods: []string{"GET", "POST", "OPTIONS"},
	}))

	return ctrlmap
}
