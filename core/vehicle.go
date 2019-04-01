package core

import "github.com/louisevanderlith/husk"

type Vehicle struct {
	VIN            *VIN
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

func (m Vehicle) Valid() (bool, error) {
	return husk.ValidateStruct(&m)
}
