package handles

import (
	"github.com/gorilla/mux"
	"github.com/louisevanderlith/kong"
	"github.com/rs/cors"
	"net/http"
)

func SetupRoutes(scrt, securityUrl, managerUrl string) http.Handler {
	r := mux.NewRouter()
	ins := kong.NewResourceInspector(http.DefaultClient, securityUrl, managerUrl)
	lkp := ins.Middleware("vin.lookup", scrt, Lookup)
	r.HandleFunc("/lookup/{vin:[A-Z0-9]+}", lkp).Methods(http.MethodGet)

	mans := ins.Middleware("vin.lookup.manufacturers", scrt, GetManufacturers)
	r.HandleFunc("/lookup/manufacturers/{year:[0-9]+}", mans).Methods(http.MethodGet)

	mdls := ins.Middleware("vin.lookup.models", scrt, GetModels)
	r.HandleFunc("/lookup/models/{year:[0-9]+}/{manufacturer:[a-zA-Z]+}", mdls).Methods(http.MethodGet)

	trms := ins.Middleware("vin.lookup.trims", scrt, GetTrims)
	r.HandleFunc("/lookup/trim/{year:[0-9]+}/{manufacturer:[a-zA-Z]+}/{model:[a-zA-Z]+}", trms).Methods(http.MethodGet)

	vald := ins.Middleware("vin.validate", scrt, Validate)
	r.HandleFunc("/validate/{vin:[A-Z0-9]+}", vald).Methods(http.MethodGet)

	viewRgn := ins.Middleware("vin.region.view", scrt, ViewRegions)
	r.HandleFunc("/regions/{key:[0-9]+\\x60[0-9]+}", viewRgn).Methods(http.MethodGet)

	getRgn := ins.Middleware("vin.region.search", scrt, GetRegions)
	r.HandleFunc("/regions", getRgn).Methods(http.MethodGet)

	srchRgn := ins.Middleware("vin.region.search", scrt, SearchRegions)
	r.HandleFunc("/regions/{pagesize:[A-Z][0-9]+}", srchRgn).Methods(http.MethodGet)
	r.HandleFunc("/regions/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", srchRgn).Methods(http.MethodGet)

	updateRgn := ins.Middleware("vin.region.update", scrt, UpdateRegion)
	r.HandleFunc("/regions/{key:[0-9]+\\x60[0-9]+}", updateRgn).Methods(http.MethodPut)

	viewAdmin := ins.Middleware("vin.admin.view", scrt, ViewAdmin)
	r.HandleFunc("/admin/{key:[0-9]+\\x60[0-9]+}", viewAdmin).Methods(http.MethodGet)

	getAdmin := ins.Middleware("vin.admin.search", scrt, GetAdmin)
	r.HandleFunc("/admin", getAdmin).Methods(http.MethodGet)

	srchAdmin := ins.Middleware("vin.admin.search", scrt, SearchAdmin)
	r.HandleFunc("/admin/{pagesize:[A-Z][0-9]+}", srchAdmin).Methods(http.MethodGet)
	r.HandleFunc("/admin/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", srchAdmin).Methods(http.MethodGet)

	lst, err := kong.Whitelist(http.DefaultClient, securityUrl, "vin.validate", scrt)

	if err != nil {
		panic(err)
	}

	corsOpts := cors.New(cors.Options{
		AllowedOrigins: lst, //you service is available and allowed for this base url
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodDelete,
			http.MethodOptions,
			http.MethodHead,
		},
		AllowCredentials: true,
		AllowedHeaders: []string{
			"*", //or you can your header key values which you are using in your application
		},
	})

	return corsOpts.Handler(r)
}
