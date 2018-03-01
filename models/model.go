package models

import "github.com/louisevanderlith/db"

type Model struct {
	db.Record
	Name         string
	Manufacturer *Manufacturer
	Engines      []*Engine
	Gearboxes    []*Gearbox
	StartYear    int
	EndYear      int
}
