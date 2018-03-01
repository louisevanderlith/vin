package models

import (
	"time"

	"github.com/louisevanderlith/db"
)

type GearboxHistory struct {
	SwapDate time.Time
	SerialNo string
	Gearbox  *Gearbox
}

type EngineHistory struct {
	db.Record
	SwapDate time.Time
	SerialNo string
	Engine   *Engine
}
