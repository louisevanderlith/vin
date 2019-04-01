package core

import "github.com/louisevanderlith/husk"

type Series struct {
	Model     Model
	Platform  Platform
	Spec      string
	StartYear int
	EndYear   int
	Vehicles  []Vehicle
}

func (m Series) Valid() (bool, error) {
	return husk.ValidateStruct(&m)
}
