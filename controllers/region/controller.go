package region

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/vin/core"
)

func Get(c *gin.Context) {
	results := core.GetAllRegions(1, 10)

	c.JSON(http.StatusOK, results)
}

// /v1/region/:key
func View(c *gin.Context) {
	k := c.Param("key")
	key, err := husk.ParseKey(k)

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	rec, err := core.GetRegion(key)

	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}

	c.JSON(http.StatusOK, rec)
}

// @router /:pagesize/:query== [get]
func Search(c *gin.Context) {
	page, size := getPageData(c.Param("pagesize"))
	results := core.GetAllRegions(page, size)

	c.JSON(http.StatusOK, results)
}

// @router /v1/region/ [put]
func Update(c *gin.Context) {
	key, err := husk.ParseKey(c.Param("key"))

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	body := &core.Region{}
	err = c.Bind(body)

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err = body.Update(key)

	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}

	c.JSON(http.StatusOK, nil)
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
