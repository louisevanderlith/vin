package core

import (
	"github.com/louisevanderlith/husk"
)

type context struct {
	VINS    husk.Tabler
	Regions husk.Tabler
}

var ctx context

func CreateContext() {
	defer seed()

	ctx = context{
		VINS:    husk.NewTable(new(VIN)),
		Regions: husk.NewTable(new(Region)),
	}
}

func Shutdown() {
	ctx.Regions.Save()
	ctx.VINS.Save()
}

func seed() {
	err := ctx.Regions.Seed("db/regions.seed.json")

	if err != nil {
		panic(err)
	}

	ctx.Regions.Save()
}
