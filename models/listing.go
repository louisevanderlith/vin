package models

import "github.com/louisevanderlith/db"

type Listing struct {
	db.Record
	UniqueVIN      string // VIN without SEQ No
	SequenceNo     int
	Series         *Series
	Colour         string
	Year           int
	AssemblyPlant  *AssemblyPlant
	EngineHistory  []*EngineHistory
	GearboxHistory []*GearboxHistory
	Alterations    []*Alteration
}
