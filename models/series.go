package models

import "github.com/louisevanderlith/db"

type Series struct {
	db.Record
	Model     *Model
	Platform  *Platform
	Spec      string
	StartYear int
	EndYear   int
	Listings  []*Listing
}
