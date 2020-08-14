package handles

import (
	"github.com/louisevanderlith/droxolite/drx"
	"github.com/louisevanderlith/droxolite/mix"
	"log"
	"net/http"

	"github.com/louisevanderlith/vin/core"
)

// @Title Validate
// @Description Attempts to validate the vin
// @Success 200 {bool} bool
// @router /:vin [get]
func Validate(w http.ResponseWriter, r *http.Request) {
	vin := drx.FindParam(r, "vin")
	err := core.ValidateVIN(vin)

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	err = mix.Write(w, mix.JSON(true))

	if err != nil {
		log.Println(err)
	}
}
