package main

import (
	"github.com/gin-gonic/gin"
	"github.com/louisevanderlith/vin/controllers/admin"
	"github.com/louisevanderlith/vin/controllers/lookup"
	"github.com/louisevanderlith/vin/controllers/region"
	"github.com/louisevanderlith/vin/controllers/validate"
	"github.com/louisevanderlith/vin/core"
)

func main() {
	core.CreateContext()
	defer core.Shutdown()

	r := gin.Default()

	// admin
	r.GET("/admin/:key", admin.View)

	// admins := r.Group("/admin")
	// admins.POST("", admin.Create)
	// admins.PUT("/:key", admin.Update)
	// admins.DELETE("/:key", admin.Delete)

	r.GET("/admin", admin.Get)
	r.GET("/admin/:pagesize/*hash", admin.Search)

	// lookup
	r.GET("/lookup/:key", lookup.Lookup)

	// lookups := r.Group("/lookup")
	// lookups.POST("", lookup.Create)
	// lookups.PUT("/:key", lookup.Update)
	// lookups.DELETE("/:key", lookup.Delete)

	// r.GET("/lookup", lookup.Get)
	// r.GET("/lookup/:pagesize/*hash", lookup.Search)

	// region
	r.GET("/region/:key", region.View)

	regions := r.Group("/region")
	// regions.POST("", region.Create)
	regions.PUT("/:key", region.Update)
	// regions.DELETE("/:key", region.Delete)

	r.GET("/region", region.Get)
	r.GET("/region/:pagesize/*hash", region.Search)

	// validate
	r.GET("/validate/:key", validate.Validate)

	// validates := r.Group("/validate")
	// validates.POST("", validate.Create)
	// validates.PUT("/:key", validate.Update)
	// validates.DELETE("/:key", validate.Delete)

	// r.GET("/validate", validate.Get)
	// r.GET("/revalidategion/:pagesize/*hash", validate.Search)

	err := r.Run(":8095")

	if err != nil {
		panic(err)
	}
}

// func main() {
// 	keyPath := os.Getenv("KEYPATH")
// 	pubName := os.Getenv("PUBLICKEY")
// 	host := os.Getenv("HOST")
// 	httpport, _ := strconv.Atoi(os.Getenv("HTTPPORT"))
// 	appName := os.Getenv("APPNAME")
// 	pubPath := path.Join(keyPath, pubName)

// 	// Register with router
// 	srv := bodies.NewService(appName, "", pubPath, host, httpport, servicetype.API)

// 	routr, err := do.GetServiceURL("", "Router.API", false)

// 	if err != nil {
// 		panic(err)
// 	}

// 	err = srv.Register(routr)

// 	if err != nil {
// 		panic(err)
// 	}

// 	poxy := resins.NewMonoEpoxy(srv, element.GetNoTheme(host, srv.ID, "none"))
// 	routers.Setup(poxy)
// 	poxy.EnableCORS(host)

// 	core.CreateContext()
// 	defer core.Shutdown()

// 	err = droxolite.Boot(poxy)

// 	if err != nil {
// 		panic(err)
// 	}
// }
