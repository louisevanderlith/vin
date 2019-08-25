package controllers

import (
	"log"
	"net/http"

	"github.com/louisevanderlith/droxolite/context"
	"github.com/louisevanderlith/vin/core"
)

type LookupController struct {
}

// @Title Validate and Deserialize
// @Description Gets the details of a VIN after validation
// @Success 200 {[]core.Profile} []core.Portfolio]
// @router /:vin [get]
func (req *LookupController) Get(ctx context.Contexer) (int, interface{}) {
	vin := ctx.FindParam("vin")
	err := core.ValidateVIN(vin)

	if err != nil {
		return http.StatusBadRequest, err
	}

	obj, err := core.BuildInfo(vin)

	if err != nil {
		log.Println("build", err)
		return http.StatusInternalServerError, err
	}

	rec, err := obj.Create()

	if err != nil {
		log.Println("create", err)
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, rec
}
