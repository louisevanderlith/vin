package controllers

import (
	"net/http"

	"github.com/louisevanderlith/droxolite/xontrols"
	"github.com/louisevanderlith/vin/core"
)

type ValidateController struct {
	xontrols.APICtrl
}

// @Title Validate
// @Description Attempts to validate the vin
// @Success 200 {bool} bool
// @router /:vin [get]
func (req *ValidateController) Get() {
	vin := req.FindParam("vin")
	err := core.ValidateVIN(vin)

	if err != nil {
		req.Serve(http.StatusBadRequest, err, nil)
		return
	}

	req.Serve(http.StatusOK, nil, true)
}
