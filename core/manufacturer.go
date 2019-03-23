package core

import "github.com/louisevanderlith/husk"

type Manufacturer struct {
	Name           string
	CommonName     string
	Description    string
	StartYear      int
	EndYear        int
	Parent         *Manufacturer
	WMIs           []*WMI
	AssemblyPlants []*AssemblyPlant
}

func (m Manufacturer) Valid() (bool, error) {
	return husk.ValidateStruct(&m)
}
