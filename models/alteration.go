package models

import (
	"time"

	"github.com/louisevanderlith/db"
)

type AlterationType = int

const (
	Maintenance AlterationType = iota
	Cosmetic
	Performance
	Utility // Canopy, Towbar, etc.
)

type Alteration struct {
	db.Record
	AlterDate   time.Time
	AlterType   AlterationType
	Code        string
	Description string
	Odometer    int64
}
