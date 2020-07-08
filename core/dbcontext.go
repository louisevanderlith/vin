package core

import (
	"github.com/louisevanderlith/husk"
)

type context struct {
	VIN     husk.Tabler
	Regions husk.Tabler
}

var ctx context

func CreateContext() {
	ctx = context{
		Regions: husk.NewTable(Region{}),
		VIN:     husk.NewTable(VIN{}),
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
