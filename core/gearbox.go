package core

import (
	"github.com/louisevanderlith/husk/validation"
)

type GearboxType = int

const (
	Manual GearboxType = iota
	Automatic
	CVT
	Sequential
)

type Gearbox struct {
	SeriesCode string
	Code       string
	Gears      int
	Type       string
	StartYear  int
	EndYear    int
}

func (m Gearbox) Valid() error {
	return validation.Struct(m)
}
