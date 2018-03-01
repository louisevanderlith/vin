package models

import (
	"github.com/louisevanderlith/db"
)

type AssemblyPlant struct {
	db.Record
	Code         string
	Manufacturer *Manufacturer
	Name         string
	Country      string
	StartYear    int
	EndYear      int
}
