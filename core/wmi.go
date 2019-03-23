package core

import "github.com/louisevanderlith/husk"

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
	Code         string
	CountryCode  string
	Manufacturer *Manufacturer
	Country      *Country
	VehicleType  VehicleType
}

func (m WMI) Valid() (bool, error) {
	return husk.ValidateStruct(&m)
}
