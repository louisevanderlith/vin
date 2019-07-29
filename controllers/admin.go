package controllers

import (
	"net/http"

	"github.com/louisevanderlith/droxolite/xontrols"
	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/vin/core"
)

type AdminController struct {
	xontrols.APICtrl
}

// /v1/vin/:key
func (req *AdminController) GetByKey() {
	k := req.FindParam("key")
	key, err := husk.ParseKey(k)

	if err != nil {
		req.Serve(http.StatusBadRequest, err, nil)
		return
	}

	rec, err := core.GetVIN(key)

	if err != nil {
		req.Serve(http.StatusNotFound, err, nil)
		return
	}

	req.Serve(http.StatusOK, nil, rec)
}

// @router /all/:pagesize [get]
func (req *AdminController) Get() {
	page, size := req.GetPageData()
	results := core.GetAllVINS(page, size)

	req.Serve(http.StatusOK, nil, results)
}
