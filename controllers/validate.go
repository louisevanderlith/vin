package controllers

import (
	"net/http"

	"github.com/louisevanderlith/mango/control"
	"github.com/louisevanderlith/vin/core"
)

type ValidateController struct {
	control.APIController
}

func NewValidateCtrl(ctrlMap *control.ControllerMap) *ValidateController {
	result := &ValidateController{}
	result.SetInstanceMap(ctrlMap)

	return result
}

// @Title Validate
// @Description Attempts to validate the vin
// @Success 200 {bool} bool
// @router /:vin [get]
func (req *ValidateController) Get() {
	vin := req.Ctx.Input.Param(":vin")
	err := core.ValidateVIN(vin)

	if err != nil {
		req.Serve(http.StatusBadRequest, err, nil)
		return
	}

	req.Serve(http.StatusOK, nil, true)
}
