package core

import "github.com/louisevanderlith/husk"

type Country struct {
	RegionCode    string
	Name          string
	StartChar     string
	EndChar       string
	Manufacturers []Manufacturer
}

func (m Country) Valid() error {
	return husk.ValidateStruct(&m)
}

func (r *Country) HasCode(regionCode string) bool {
	s := getCharWeight(r.StartChar)
	e := getCharWeight(r.EndChar)
	v := getCharWeight(regionCode)

	return s <= v && v <= e
}
