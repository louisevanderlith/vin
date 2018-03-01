package models

import (
	"github.com/louisevanderlith/db"
)

type VehicleType = int

const (
	PassengerCar VehicleType = iota
	Motorcycle
	Truck
	MPV
	Trailer
	LSV // Low speed vehicle
	ATV
	Incomplete
)

type WMI struct {
	db.Record
	Code         string
	CountryCode  string
	Manufacturer *Manufacturer
	Country      *Country
	VehicleType  VehicleType
}
