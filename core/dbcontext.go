package core

import (
	"github.com/louisevanderlith/husk"
)

type context struct {
	VINS    husk.Tabler
	Regions husk.Tabler
}

var ctx context

func init() {
	defer seed();

	ctx = context{
		VINS:    husk.NewTable(new(VIN)),
		Regions: husk.NewTable(new(Region)),
	}
}

func seed() {
	//conf/regions.seed.json
	
}