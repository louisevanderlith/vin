package handles

import (
	"github.com/gorilla/mux"
	"github.com/louisevanderlith/kong"
	"github.com/rs/cors"
	"net/http"
)

func SetupRoutes(scrt, secureUrl string) http.Handler {
	r := mux.NewRouter()

	mans := kong.ResourceMiddleware("vin.lookup.manufacturers", scrt, secureUrl, GetManufacturers)
	r.HandleFunc("/lookup/manufacturers/{year:[0-9]+}", mans).Methods(http.MethodGet)

	mdls := kong.ResourceMiddleware("vin.lookup.models", scrt, secureUrl, GetModels)
	r.HandleFunc("/lookup/models/{year:[0-9]+}/{manufacturer:[a-zA-Z]+}", mdls).Methods(http.MethodGet)

	trms := kong.ResourceMiddleware("vin.lookup.trims", scrt, secureUrl, GetTrims)
	r.HandleFunc("/lookup/trim/{year:[0-9]+}/{manufacturer:[a-zA-Z]+}/{model:[a-zA-Z]+}", trms).Methods(http.MethodGet)

	vald := kong.ResourceMiddleware("vin.validate", scrt, secureUrl, Validate)
	r.HandleFunc("/validate/{vin}", vald).Methods(http.MethodGet)

	/*
		admCtrl := &handles.Admin{}
			regnCtrl := &handles.Regions{}
			e.JoinBundle("/", roletype.Admin, mix.JSON, admCtrl, regnCtrl)
			e.JoinPath(e.Router().(*mux.Router), "/lookup/{vin}", "Find VIN", http.MethodGet, roletype.User, mix.JSON, handles.Lookup)
			e.JoinPath(e.Router().(*mux.Router), "/validate/{vin}", "Find VIN", http.MethodGet, roletype.User, mix.JSON, handles.Validate)
	*/

	/*
		{
		    "Name": "vin.lookup",
		    "DisplayName": "VIN Lookup",
		    "Secret": "secret"
		  },
		  {
		    "Name": "vin.validate",
		    "DisplayName": "VIN Validation",
		    "Secret": "secret"
		  },
		  {
		    "Name": "vin.region.view",
		    "DisplayName": "View vin Region",
		    "Secret": "secret"
		  },
		  {
		    "Name": "vin.region.create",
		    "DisplayName": "Create vin Region",
		    "Secret": "secret"
		  },
		  {
		    "Name": "vin.region.update",
		    "DisplayName": "Update vin Region",
		    "Secret": "secret"
		  },
		  {
		    "Name": "vin.region.search",
		    "DisplayName": "Search vin Regions",
		    "Secret": "secret"
		  },
		  {
		    "Name": "vin.admin.view",
		    "DisplayName": "View vin Region",
		    "Secret": "secret"
		  },
		  {
		    "Name": "vin.admin.search",
		    "DisplayName": "Search vin Regions",
		    "Secret": "secret"
		  },
	*/
	lst, err := kong.Whitelist(http.DefaultClient, secureUrl, "comment.messages.view", scrt)

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
