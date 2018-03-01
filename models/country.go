package models

import (
	"github.com/louisevanderlith/db"
)

type Country struct {
	db.Record
	ContinentCode string
	Name          string
	StartChar     string
	EndChar       string
	Continent     *Continent
	WMIs          []*WMI
}
