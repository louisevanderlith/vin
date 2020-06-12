package handles

import (
	"github.com/louisevanderlith/droxolite/mix"
	"log"
	"net/http"

	"github.com/louisevanderlith/droxolite/context"
	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/vin/core"
)

func GetRegions(w http.ResponseWriter, r *http.Request) {
	ctx := context.New(w, r)
	results, err := core.GetAllRegions(1, 10)

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	err = ctx.Serve(http.StatusOK, mix.JSON(results))

	if err != nil {
		log.Println(err)
	}
}

// /v1/region/:key
func ViewRegions(w http.ResponseWriter, r *http.Request) {
	ctx := context.New(w, r)
	k := ctx.FindParam("key")
	key, err := husk.ParseKey(k)

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	rec, err := core.GetRegion(key)

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusNotFound)
		return
	}

	err = ctx.Serve(http.StatusOK, mix.JSON(rec))

	if err != nil {
		log.Println(err)
	}
}

// @router /:pagesize/:query== [get]
func SearchRegions(w http.ResponseWriter, r *http.Request) {
	ctx := context.New(w, r)
	page, size := ctx.GetPageData()
	results, err := core.GetAllRegions(page, size)

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	err = ctx.Serve(http.StatusOK, mix.JSON(results))

	if err != nil {
		log.Println(err)
	}
}

// @router /v1/region/ [put]
func UpdateRegion(w http.ResponseWriter, r *http.Request) {
	ctx := context.New(w, r)
	key, err := husk.ParseKey(ctx.FindParam("key"))

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	body := &core.Region{}
	err = ctx.Body(body)

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	err = body.Update(key)

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusNotFound)
		return
	}

	err = ctx.Serve(http.StatusOK, mix.JSON(nil))

	if err != nil {
		log.Println(err)
	}
}
