package controllers

import (
	"net/http"

	"github.com/louisevanderlith/droxolite/context"
	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/vin/core"
)

type Admin struct {
}

func (x *Admin) Get(ctx context.Requester) (int, interface{}) {
	results := core.GetAllVINS(1, 10)

	return http.StatusOK, results
}

// /v1/vin/:key
func (req *Admin) View(ctx context.Requester) (int, interface{}) {
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
func (req *Admin) Search(ctx context.Requester) (int, interface{}) {
	page, size := ctx.GetPageData()
	results := core.GetAllVINS(page, size)

	return http.StatusOK, results
}
