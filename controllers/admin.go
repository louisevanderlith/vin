package controllers

import (
	"net/http"

	"github.com/louisevanderlith/droxolite/context"
	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/vin/core"
)

type Admin struct {
}

func (x *Admin) Get(c *gin.Context) {
	results := core.GetAllVINS(1, 10)

	return http.StatusOK, results
}

// /v1/vin/:key
func (req *Admin) View(c *gin.Context) {
	k := c.Param("key")
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
func (req *Admin) Search(c *gin.Context) {
	page, size := ctx.GetPageData()
	results := core.GetAllVINS(page, size)

	return http.StatusOK, results
}
