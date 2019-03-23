package core

import "github.com/louisevanderlith/husk"

type AssemblyPlant struct {
	Code         string
	Manufacturer *Manufacturer
	Name         string
	Country      string
	StartYear    int
	EndYear      int
}

func (m AssemblyPlant) Valid() (bool, error) {
	return husk.ValidateStruct(&m)
}
