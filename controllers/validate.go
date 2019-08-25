package controllers

import (
	"net/http"

	"github.com/louisevanderlith/droxolite/context"
	"github.com/louisevanderlith/vin/core"
)

type ValidateController struct {
}

// @Title Validate
// @Description Attempts to validate the vin
// @Success 200 {bool} bool
// @router /:vin [get]
func (req *ValidateController) Get(ctx context.Contexer) (int, interface{}) {
	vin := ctx.FindParam("vin")
	err := core.ValidateVIN(vin)

	if err != nil {
		return http.StatusBadRequest, err
	}

	return http.StatusOK, true
}
