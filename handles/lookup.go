package handles

import (
	"github.com/louisevanderlith/droxolite/mix"
	"log"
	"net/http"
	"strconv"

	"github.com/louisevanderlith/droxolite/context"
	"github.com/louisevanderlith/vin/core"
)

// @Title Validate and Deserialize
// @Description Gets the details of a VIN after validation
// @Success 200 {[]core.Profile} []core.Portfolio]
// @router /:vin [get]
func Lookup(w http.ResponseWriter, r *http.Request) {
	ctx := context.New(w, r)
	vin := ctx.FindParam("vin")
	err := core.ValidateVIN(vin)

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	obj, err := core.BuildInfo(vin)

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	rec, err := obj.Create()

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	err = ctx.Serve(http.StatusOK, mix.JSON(rec))

	if err != nil {
		log.Println(err)
	}
}

func GetManufacturers(w http.ResponseWriter, r *http.Request) {
	ctx := context.New(w, r)

	year, _ := strconv.Atoi(ctx.FindParam("year"))
	result, err := core.GetManufacturers(year)

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	var lst []string

	for name, _ := range result {
		lst = append(lst, name)
	}

	err = ctx.Serve(http.StatusOK, mix.JSON(lst))

	if err != nil {
		log.Println(err)
	}
}

func GetModels(w http.ResponseWriter, r *http.Request) {
	ctx := context.New(w, r)

	year, _ := strconv.Atoi(ctx.FindParam("year"))
	man := ctx.FindParam("manufacturer")

	result, err := core.GetModels(year, man)

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	var lst []string

	for name, _ := range result {
		lst = append(lst, name)
	}

	err = ctx.Serve(http.StatusOK, mix.JSON(lst))

	if err != nil {
		log.Println(err)
	}
}

func GetTrims(w http.ResponseWriter, r *http.Request) {
	ctx := context.New(w, r)

	year, _ := strconv.Atoi(ctx.FindParam("year"))
	man := ctx.FindParam("manufacturer")
	mdl := ctx.FindParam("model")

	result, err := core.GetTrims(year, man, mdl)

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	var lst []string

	for name, _ := range result {
		lst = append(lst, name)
	}

	err = ctx.Serve(http.StatusOK, mix.JSON(lst))

	if err != nil {
		log.Println(err)
	}
}
