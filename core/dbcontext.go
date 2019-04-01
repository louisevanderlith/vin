package core

import (
	"github.com/louisevanderlith/husk"
)

type context struct {
	Vehicles husk.Tabler
	Regions  husk.Tabler
}

var ctx context

func CreateContext() {
	defer seed()

	ctx = context{
		Vehicles: husk.NewTable(new(Vehicle)),
		Regions:  husk.NewTable(new(Region)),
	}
}

func Shutdown() {
	ctx.Regions.Save()
	ctx.Vehicles.Save()
}

func seed() {
	err := ctx.Regions.Seed("db/regions.seed.json")

	if err != nil {
		panic(err)
	}

	ctx.Regions.Save()
}
