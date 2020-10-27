package core

import (
	"github.com/louisevanderlith/husk/validation"
)

type EngineLayout = int

const (
	Inline EngineLayout = iota
	V
	Rotary
	W
	Boxer
)

type FuelType = int

const (
	Petrol FuelType = iota
	Diesel
	Hybrid
	Electric
	LPG
)

type Induction = int

const (
	NA Induction = iota
	Turbo
	Supercharger
)

type Engine struct {
	Family            string
	Series            string
	Code              string
	Displacement      int
	FuelType          string
	Layout            string
	Cylinders         int
	Valvetrain        string
	ValvesPerCylinder int
	PowerKW           int
	PowerAt           int
	TorqueNm          int
	TorqueAt          int
	Induction         string
	StartYear         int
	EndYear           int
}

func (m Engine) Valid() error {
	return validation.Struct(m)
}
