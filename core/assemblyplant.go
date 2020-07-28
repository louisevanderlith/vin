package core

import "github.com/louisevanderlith/husk"

type AssemblyPlant struct {
	Code      string
	Name      string
	Country   string
	StartYear int
	EndYear   int
	Series    []Series
}

func (m AssemblyPlant) Valid() error {
	return husk.ValidateStruct(m)
}
