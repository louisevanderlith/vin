package models

import "github.com/louisevanderlith/db"

type GearboxType = int

const (
	Manual GearboxType = iota
	Automatic
	CVT
	Sequential
)

type Gearbox struct {
	db.Record
	SeriesCode string
	Code       string
	Gears      int
	Type       GearboxType
	Models     []*Model
	StartYear  int
	EndYear    int
}
