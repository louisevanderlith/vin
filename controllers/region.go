package controllers

import (
	"net/http"

	"github.com/louisevanderlith/droxolite/context"
	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/vin/core"
)

type Regions struct {
}

func (req *Regions) Get(c *gin.Context) {
	results := core.GetAllRegions(1, 10)

	return http.StatusOK, results
}

// /v1/region/:key
func (req *Regions) View(c *gin.Context) {
	k := c.Param("key")
	key, err := husk.ParseKey(k)

	if err != nil {
		return http.StatusBadRequest, err
	}

	rec, err := core.GetRegion(key)

	if err != nil {
		return http.StatusNotFound, err
	}

	return http.StatusOK, rec
}

// @router /:pagesize/:query== [get]
func (req *Regions) Search(c *gin.Context) {
	page, size := ctx.GetPageData()
	results := core.GetAllRegions(page, size)

	return http.StatusOK, results
}

// @router /v1/region/ [put]
func (req *Regions) Update(c *gin.Context) {
	key, err := husk.ParseKey(c.Param("key"))

	if err != nil {
		return http.StatusBadRequest, err
	}

	body := &core.Region{}
	err = ctx.Body(body)

	if err != nil {
		return http.StatusBadRequest, err
	}

	err = body.Update(key)

	if err != nil {
		return http.StatusNotFound, err
	}

	return http.StatusOK, nil
}
