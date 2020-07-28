package core

import "github.com/louisevanderlith/husk"

type DriveLayout = int

const (
	FrontFront DriveLayout = iota
	FrontRear
	FrontFour
	MidFront
	MidRear
	MidFour
	RearFront
	RearRear
	RearFour
)

type Platform struct {
	Code        string
	Engine      Engine
	Gearbox     Gearbox
	Body        Body
	DriveLayout string
	StartYear   int
	EndYear     int
}

func (m Platform) Valid() error {
	return husk.ValidateStruct(m)
}

/*

*/