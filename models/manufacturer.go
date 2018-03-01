package models

import "github.com/louisevanderlith/db"

type Manufacturer struct {
	db.Record
	Name           string
	CommonName     string
	Description    string
	StartYear      int
	EndYear        int
	Parent         *Manufacturer
	WMIs           []*WMI
	AssemblyPlants []*AssemblyPlant
}
