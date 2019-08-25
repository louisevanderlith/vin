package controllers

import (
	"net/http"

	"github.com/louisevanderlith/droxolite/context"
	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/vin/core"
)

type AdminController struct {
}

// /v1/vin/:key
func (req *AdminController) GetByKey(ctx context.Contexer) (int, interface{}) {
	k := ctx.FindParam("key")
	key, err := husk.ParseKey(k)

	if err != nil {
		return http.StatusBadRequest, err
	}

	rec, err := core.GetVIN(key)

	if err != nil {
		return http.StatusNotFound, err
	}

	return http.StatusOK, rec
}

// @router /all/:pagesize [get]
func (req *AdminController) Get(ctx context.Contexer) (int, interface{}) {
	page, size := ctx.GetPageData()
	results := core.GetAllVINS(page, size)

	return http.StatusOK, results
}
