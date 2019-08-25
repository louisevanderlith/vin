package routers

import (
	"github.com/louisevanderlith/droxolite/mix"
	"github.com/louisevanderlith/droxolite/resins"
	"github.com/louisevanderlith/droxolite/roletype"
	"github.com/louisevanderlith/droxolite/routing"

	"github.com/louisevanderlith/vin/controllers"
)

func Setup(poxy resins.Epoxi) {
	//Admin
	admCtrl := &controllers.AdminController{}
	admGroup := routing.NewRouteGroup("admin", mix.JSON)
	admGroup.AddRoute("VIN by Key", "/{key:[0-9]+\x60[0-9]+}", "GET", roletype.Admin, admCtrl.GetByKey)
	admGroup.AddRoute("All VIN Numbers", "/all/{pagesize:[A-Z][0-9]+}", "GET", roletype.Admin, admCtrl.Get)
	poxy.AddGroup(admGroup)

	//Region
	regnCtrl := &controllers.RegionController{}
	regnGroup := routing.NewRouteGroup("region", mix.JSON)
	regnGroup.AddRoute("Region by Key", "/{key:[0-9]+\x60[0-9]+}", "GET", roletype.User, regnCtrl.GetByKey)
	regnGroup.AddRoute("All Regions", "/all/{pagesize:[A-Z][0-9]+}", "GET", roletype.User, regnCtrl.Get)
	poxy.AddGroup(regnGroup)

	//Lookup
	lookCtrl := &controllers.LookupController{}
	lookGroup := routing.NewRouteGroup("lookup", mix.JSON)
	lookGroup.AddRoute("Find VIN", "/{vin}", "GET", roletype.User, lookCtrl.Get)
	poxy.AddGroup(lookGroup)

	//Validate
	valCtrl := &controllers.ValidateController{}
	valGroup := routing.NewRouteGroup("validate", mix.JSON)
	valGroup.AddRoute("Validate VIN", "/{vin}", "GET", roletype.User, valCtrl.Get)
	poxy.AddGroup(valGroup)
	/*ctrlmap := EnableFilter(s, host)

	adminCtrl := controllers.NewAdminCtrl(ctrlmap)
	beego.Router("/v1/admin/:key", adminCtrl, "get:GetByKey")
	beego.Router("/v1/admin/all/:pagesize", adminCtrl, "get:Get")

	regionCtrl := controllers.NewRegionCtrl(ctrlmap)
	beego.Router("/v1/region/:key", regionCtrl, "get:GetByKey")
	beego.Router("/v1/region/all/:pagesize", regionCtrl, "get:Get")

	beego.Router("/v1/lookup/:vin", controllers.NewLookupCtrl(ctrlmap), "get:Get")
	beego.Router("/v1/validate/:vin", controllers.NewValidateCtrl(ctrlmap), "get:Get")*/
}

/*
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
	ctrlmap.Add("/v1/region", adminMap)

	beego.InsertFilter("/v1/*", beego.BeforeRouter, ctrlmap.FilterAPI, false)
	allowed := fmt.Sprintf("https://*%s", strings.TrimSuffix(host, "/"))

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins: []string{allowed},
		AllowMethods: []string{"GET", "POST", "OPTIONS"},
	}))

	return ctrlmap
}*/
