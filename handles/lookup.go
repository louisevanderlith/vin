package handles

import (
	"github.com/louisevanderlith/droxolite/drx"
	"github.com/louisevanderlith/droxolite/mix"
	"log"
	"net/http"
	"strconv"

	"github.com/louisevanderlith/vin/core"
)

// @Title Validate and Deserialize
// @Description Gets the details of a VIN after validation
// @Success 200 {[]core.Profile} []core.Portfolio]
// @router /:vin [get]
func Lookup(w http.ResponseWriter, r *http.Request) {
	vin := drx.FindParam(r, "vin")
	err := core.Context().ValidateVIN(vin)

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	obj, err := core.Context().BuildInfo(vin)

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	item, err := core.Context().FindVIN(vin)

	if err != nil && item == nil {
		log.Println("Find Error", err)
		err = core.Context().CreateVIN(*obj)
	}

	if err != nil {
		log.Println("Create Error", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	err = mix.Write(w, mix.JSON(item))

	if err != nil {
		log.Println(err)
	}
}

func GetManufacturers(w http.ResponseWriter, r *http.Request) {
	year, _ := strconv.Atoi(drx.FindParam(r, "year"))
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

	err = mix.Write(w, mix.JSON(lst))

	if err != nil {
		log.Println(err)
	}
}

func GetModels(w http.ResponseWriter, r *http.Request) {
	year, _ := strconv.Atoi(drx.FindParam(r, "year"))
	man := drx.FindParam(r, "manufacturer")

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

	err = mix.Write(w, mix.JSON(lst))

	if err != nil {
		log.Println(err)
	}
}

func GetTrims(w http.ResponseWriter, r *http.Request) {
	year, _ := strconv.Atoi(drx.FindParam(r, "year"))
	man := drx.FindParam(r, "manufacturer")
	mdl := drx.FindParam(r, "model")

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

	err = mix.Write(w, mix.JSON(lst))

	if err != nil {
		log.Println(err)
	}
}
