package core

import "github.com/louisevanderlith/husk"

type Country struct {
	RegionCode string
	Name       string
	StartChar  string
	EndChar    string
	Region     *Region
	WMIs       []*WMI
}

func (m Country) Valid() (bool, error) {
	return husk.ValidateStruct(&m)
}
