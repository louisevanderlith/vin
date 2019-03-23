package core

import "github.com/louisevanderlith/husk"

type Model struct {
	Name         string
	Manufacturer *Manufacturer
	Engines      []*Engine
	Gearboxes    []*Gearbox
	StartYear    int
	EndYear      int
}

func (m Model) Valid() (bool, error) {
	return husk.ValidateStruct(&m)
}
