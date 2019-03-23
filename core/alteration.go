package core

import (
	"time"

	"github.com/louisevanderlith/husk"
)

type AlterationType = int

const (
	Maintenance AlterationType = iota
	Cosmetic
	Performance
	Utility // Canopy, Towbar, etc.
)

type Alteration struct {
	AlterDate   time.Time
	AlterType   AlterationType
	Code        string
	Description string
	Odometer    int64
}

func (m Alteration) Valid() (bool, error) {
	return husk.ValidateStruct(&m)
}
