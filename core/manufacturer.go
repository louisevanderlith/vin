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

type Manufacturer struct {
	WMICode        string
	Name           string
	Description    string
	VehicleType    VehicleType
	AssemblyPlants []AssemblyPlant
}

func (m Manufacturer) Valid() (bool, error) {
	return husk.ValidateStruct(&m)
}
