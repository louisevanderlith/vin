package controllers

import (
	"net/http"

	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/mango/control"
	"github.com/louisevanderlith/vin/core"
)

type RegionController struct {
	control.APIController
}

func NewRegionCtrl(ctrlmap *control.ControllerMap) *RegionController {
	result := &RegionController{}
	result.SetInstanceMap(ctrlmap)

	return result
}

// /v1/region/:key
func (req *RegionController) GetByKey() {
	k := req.Ctx.Input.Param(":key")
	key, err := husk.ParseKey(k)

	if err != nil {
		req.Serve(http.StatusBadRequest, err, nil)
		return
	}

	rec, err := core.GetRegion(key)

	if err != nil {
		req.Serve(http.StatusNotFound, err, nil)
		return
	}

	req.Serve(http.StatusOK, nil, rec)
}

// @router /all/:pagesize [get]
func (req *RegionController) Get() {
	page, size := req.GetPageData()
	results := core.GetAllRegions(page, size)

	req.Serve(http.StatusOK, nil, results)
}

// @router /v1/region/ [put]
func (req *RegionController) Put() {
	body := &core.Region{}
	key, err := req.GetKeyedRequest(body)

	if err != nil {
		req.Serve(http.StatusBadRequest, err, nil)
		return
	}

	err = body.Update(key)

	if err != nil {
		req.Serve(http.StatusNotFound, err, nil)
		return
	}

	req.Serve(http.StatusOK, nil, nil)
}
