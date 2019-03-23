package core

import (
	"github.com/louisevanderlith/husk"
)

type context struct {
	VINS husk.Tabler
}

var ctx context

func init() {
	//defer seed();

	ctx = context{
		VINS: husk.NewTable(new(VIN)),
	}
}
