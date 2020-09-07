package core

import (
	"github.com/louisevanderlith/husk/validation"
)

type Series struct {
	Platform  Platform
	Spec      string
	StartYear int
	EndYear   int
}

func (m Series) Valid() error {
	return validation.Struct(m)
}
