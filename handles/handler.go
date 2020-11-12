package handles

import (
	"github.com/gorilla/mux"
	"github.com/louisevanderlith/droxolite/open"
	"github.com/rs/cors"
	"net/http"
)

func SetupRoutes(issuer, audience string) http.Handler {
	r := mux.NewRouter()
	mw := open.BearerMiddleware(audience, issuer)
	r.Handle("/lookup/{vin:[A-Z0-9]+}", mw.Handler(http.HandlerFunc(Lookup))).Methods(http.MethodGet)
	r.Handle("/lookup/manufacturers/{year:[0-9]+}", mw.Handler(http.HandlerFunc(GetManufacturers))).Methods(http.MethodGet)
	r.Handle("/lookup/models/{year:[0-9]+}/{manufacturer:[a-zA-Z]+}", mw.Handler(http.HandlerFunc(GetModels))).Methods(http.MethodGet)
	r.Handle("/lookup/trim/{year:[0-9]+}/{manufacturer:[a-zA-Z]+}/{model:[a-zA-Z]+}", mw.Handler(http.HandlerFunc(GetTrims))).Methods(http.MethodGet)
	r.Handle("/validate/{vin:[A-Z0-9]+}", mw.Handler(http.HandlerFunc(Validate))).Methods(http.MethodGet)
	r.Handle("/regions/{key:[0-9]+\\x60[0-9]+}", mw.Handler(http.HandlerFunc(ViewRegions))).Methods(http.MethodGet)
	r.Handle("/regions", mw.Handler(http.HandlerFunc(GetRegions))).Methods(http.MethodGet)
	r.Handle("/regions/{pagesize:[A-Z][0-9]+}", mw.Handler(http.HandlerFunc(SearchRegions))).Methods(http.MethodGet)
	r.Handle("/regions/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", mw.Handler(http.HandlerFunc(SearchRegions))).Methods(http.MethodGet)
	r.Handle("/regions/{key:[0-9]+\\x60[0-9]+}", mw.Handler(http.HandlerFunc(UpdateRegion))).Methods(http.MethodPut)
	r.Handle("/admin/{key:[0-9]+\\x60[0-9]+}", mw.Handler(http.HandlerFunc(ViewAdmin))).Methods(http.MethodGet)
	r.Handle("/admin", mw.Handler(http.HandlerFunc(GetAdmin))).Methods(http.MethodGet)
	r.Handle("/admin/{pagesize:[A-Z][0-9]+}", mw.Handler(http.HandlerFunc(SearchAdmin))).Methods(http.MethodGet)
	r.Handle("/admin/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", mw.Handler(http.HandlerFunc(SearchAdmin))).Methods(http.MethodGet)

	corsOpts := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, //you service is available and allowed for this base url
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
