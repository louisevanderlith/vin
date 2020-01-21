package validate

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/louisevanderlith/vin/core"
)

// @Title Validate
// @Description Attempts to validate the vin
// @Success 200 {bool} bool
// @router /:vin [get]
func Validate(c *gin.Context) {
	vin := c.Param("vin")
	err := core.ValidateVIN(vin)

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	c.JSON(http.StatusOK, true)
}
