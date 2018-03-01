package models

import "github.com/louisevanderlith/db"

type BodyLayout = int

const (
	Sedan BodyLayout = iota
	Coupe
	Hatchback
	Van
	PickupTruck
	StationWagon
	Convertible
	SUV
	Fastback
)

type Body struct {
	db.Record
	Code      string
	Layout    BodyLayout
	Doors     int
	StartYear int
	EndYear   int
}
