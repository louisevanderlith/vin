package core

import "github.com/louisevanderlith/husk"

type Series struct {
	Platform  Platform
	Spec      string
	StartYear int
	EndYear   int
}

func (m Series) Valid() (bool, error) {
	return husk.ValidateStruct(&m)
}
