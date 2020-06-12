package core

import (
	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/husk/serials"
)

type context struct {
	VIN     husk.Tabler
	Regions husk.Tabler
}

var ctx context

func CreateContext() {
	ctx = context{
		Regions: husk.NewTable(Region{}, serials.GobSerial{}),
		VIN:     husk.NewTable(VIN{}, serials.GobSerial{}),
	}
	
	seed()
}

func Shutdown() {
	ctx.Regions.Save()
	ctx.VIN.Save()
}

func seed() {
	err := ctx.Regions.Seed("db/regions.seed.json")

	if err != nil {
		panic(err)
	}

	ctx.Regions.Save()
}
