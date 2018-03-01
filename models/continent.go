package models

import "github.com/louisevanderlith/db"

type Continent struct {
	db.Record
	Name      string
	StartChar string
	EndChar   string
	Countries []*Country
}
