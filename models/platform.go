package models

import "github.com/louisevanderlith/db"

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
	db.Record
	Code        string
	Engine      *Engine
	Gearbox     *Gearbox
	Body        *Body
	DriveLayout DriveLayout
	StartYear   int
	EndYear     int
}
