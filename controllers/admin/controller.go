package admin

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/vin/core"
)

func Get(c *gin.Context) {
	results := core.GetAllVINS(1, 10)

	c.JSON(http.StatusOK, results)
}

// /v1/vin/:key
func View(c *gin.Context) {
	k := c.Param("key")
	key, err := husk.ParseKey(k)

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	rec, err := core.GetVIN(key)

	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}

	c.JSON(http.StatusOK, rec)
}

// @router /all/:pagesize [get]
func Search(c *gin.Context) {
	page, size := getPageData(c.Param("pagesize"))
	results := core.GetAllVINS(page, size)

	c.JSON(http.StatusOK, results)
}

func getPageData(pageData string) (int, int) {
	defaultPage := 1
	defaultSize := 10

	if len(pageData) < 2 {
		return defaultPage, defaultSize
	}

	pChar := []rune(pageData[:1])

	if len(pChar) != 1 {
		return defaultPage, defaultSize
	}

	page := int(pChar[0]) % 32
	pageSize, err := strconv.Atoi(pageData[1:])

	if err != nil {
		return defaultPage, defaultSize
	}

	return page, pageSize
}
