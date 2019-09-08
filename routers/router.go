package routers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/louisevanderlith/droxolite/mix"
	"github.com/louisevanderlith/droxolite/resins"
	"github.com/louisevanderlith/droxolite/roletype"

	"github.com/louisevanderlith/vin/controllers"
)

func Setup(e resins.Epoxi) {
	admCtrl := &controllers.Admin{}
	regnCtrl := &controllers.Regions{}
	e.JoinBundle("/", roletype.Admin, mix.JSON, admCtrl, regnCtrl)
	e.JoinPath(e.Router().(*mux.Router), "/lookup/{vin}", "Find VIN", http.MethodGet, roletype.User, mix.JSON, controllers.Lookup)
	e.JoinPath(e.Router().(*mux.Router), "/validate/{vin}", "Find VIN", http.MethodGet, roletype.User, mix.JSON, controllers.Validate)
}
