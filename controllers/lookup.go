package controllers

import (
	"github.com/louisevanderlith/mango/control"
	"github.com/louisevanderlith/vin/core"
)

type LookupController struct {
	control.APIController
}

func NewLookupCtrl(ctrlMap *control.ControllerMap) *LookupController {
	result := &LookupController{}
	result.SetInstanceMap(ctrlMap)

	return result
}

// @Title Validate and Deserialize
// @Description Gets the details of a VIN after validation
// @Success 200 {[]core.Profile} []core.Portfolio]
// @router /:vin [get]
func (req *LookupController) Get() {
	vin := req.Ctx.Input.Param(":vin")
	err := core.ValidateVIN(vin)

	if err != nil {
		req.Serve(nil, err)
		return
	}

	results, err := core.doesVINExist(vin)

	req.Serve(results, err)
}
