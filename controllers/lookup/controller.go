package lookup

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/louisevanderlith/vin/core"
)

// @Title Validate and Deserialize
// @Description Gets the details of a VIN after validation
// @Success 200 {[]core.Profile} []core.Portfolio]
// @router /:vin [get]
func Lookup(c *gin.Context) {
	vin := c.Param("vin")
	err := core.ValidateVIN(vin)

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	obj, err := core.BuildInfo(vin)

	if err != nil {
		log.Println("build", err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	rec, err := obj.Create()

	if err != nil {
		log.Println("create", err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, rec)
}
