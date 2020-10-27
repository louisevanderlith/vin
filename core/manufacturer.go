package core

import (
	"github.com/louisevanderlith/husk/validation"
)

type VehicleType int

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

var vehTypes = [...]string{
	"PassengerCar",
	"Motorcycle",
	"Truck",
	"MPV",
	"Trailer",
	"LSV",
	"ATV",
	"Incomplete"}

func (s VehicleType) String() string {
	return vehTypes[s]
}

type Manufacturer struct {
	WMICode        string
	Name           string
	Description    string
	VehicleType    VehicleType
	AssemblyPlants []AssemblyPlant
}

func (m Manufacturer) Valid() error {
	return validation.Struct(m)
}
