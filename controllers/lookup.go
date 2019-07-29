package controllers

import (
	"log"
	"net/http"

	"github.com/louisevanderlith/droxolite/xontrols"
	"github.com/louisevanderlith/vin/core"
)

type LookupController struct {
	xontrols.APICtrl
}

// @Title Validate and Deserialize
// @Description Gets the details of a VIN after validation
// @Success 200 {[]core.Profile} []core.Portfolio]
// @router /:vin [get]
func (req *LookupController) Get() {
	vin := req.FindParam("vin")
	err := core.ValidateVIN(vin)

	if err != nil {
		req.Serve(http.StatusBadRequest, err, nil)
		return
	}

	obj, err := core.BuildInfo(vin)

	if err != nil {
		log.Println("build", err)
		req.Serve(http.StatusInternalServerError, err, nil)
		return
	}

	rec, err := obj.Create()

	if err != nil {
		log.Println("create", err)
		req.Serve(http.StatusInternalServerError, err, nil)
		return
	}

	req.Serve(http.StatusOK, nil, rec)
}
