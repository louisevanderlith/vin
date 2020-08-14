package handles

import (
	"github.com/gorilla/mux"
	"github.com/louisevanderlith/kong"
	"github.com/rs/cors"
	"net/http"
)

func SetupRoutes(scrt, securityUrl, managerUrl string) http.Handler {
	r := mux.NewRouter()

	lkp := kong.ResourceMiddleware(http.DefaultClient, "vin.lookup", scrt, securityUrl, managerUrl, Lookup)
	r.HandleFunc("/lookup/{vin:[A-Z0-9]+}", lkp).Methods(http.MethodGet)

	mans := kong.ResourceMiddleware(http.DefaultClient, "vin.lookup.manufacturers", scrt, securityUrl, managerUrl, GetManufacturers)
	r.HandleFunc("/lookup/manufacturers/{year:[0-9]+}", mans).Methods(http.MethodGet)

	mdls := kong.ResourceMiddleware(http.DefaultClient, "vin.lookup.models", scrt, securityUrl, managerUrl, GetModels)
	r.HandleFunc("/lookup/models/{year:[0-9]+}/{manufacturer:[a-zA-Z]+}", mdls).Methods(http.MethodGet)

	trms := kong.ResourceMiddleware(http.DefaultClient, "vin.lookup.trims", scrt, securityUrl, managerUrl, GetTrims)
	r.HandleFunc("/lookup/trim/{year:[0-9]+}/{manufacturer:[a-zA-Z]+}/{model:[a-zA-Z]+}", trms).Methods(http.MethodGet)

	vald := kong.ResourceMiddleware(http.DefaultClient, "vin.validate", scrt, securityUrl, managerUrl, Validate)
	r.HandleFunc("/validate/{vin:[A-Z0-9]+}", vald).Methods(http.MethodGet)

	viewRgn := kong.ResourceMiddleware(http.DefaultClient, "vin.region.view", scrt, securityUrl, managerUrl, ViewRegions)
	r.HandleFunc("/regions/{key:[0-9]+\\x60[0-9]+}", viewRgn).Methods(http.MethodGet)

	getRgn := kong.ResourceMiddleware(http.DefaultClient, "vin.region.search", scrt, securityUrl, managerUrl, GetRegions)
	r.HandleFunc("/regions", getRgn).Methods(http.MethodGet)

	srchRgn := kong.ResourceMiddleware(http.DefaultClient, "vin.region.search", scrt, securityUrl, managerUrl, SearchRegions)
	r.HandleFunc("/regions/{pagesize:[A-Z][0-9]+}", srchRgn).Methods(http.MethodGet)
	r.HandleFunc("/regions/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", srchRgn).Methods(http.MethodGet)

	updateRgn := kong.ResourceMiddleware(http.DefaultClient, "vin.region.update", scrt, securityUrl, managerUrl, UpdateRegion)
	r.HandleFunc("/regions/{key:[0-9]+\\x60[0-9]+}", updateRgn).Methods(http.MethodPut)

	viewAdmin := kong.ResourceMiddleware(http.DefaultClient, "vin.admin.view", scrt, securityUrl, managerUrl, ViewAdmin)
	r.HandleFunc("/admin/{key:[0-9]+\\x60[0-9]+}", viewAdmin).Methods(http.MethodGet)

	getAdmin := kong.ResourceMiddleware(http.DefaultClient, "vin.admin.search", scrt, securityUrl, managerUrl, GetAdmin)
	r.HandleFunc("/admin", getAdmin).Methods(http.MethodGet)

	srchAdmin := kong.ResourceMiddleware(http.DefaultClient, "vin.admin.search", scrt, securityUrl, managerUrl, SearchAdmin)
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
