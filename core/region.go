package core

import "github.com/louisevanderlith/husk"

type Region struct {
	Name      string
	StartChar string
	EndChar   string
	Countries []*Country
}

func (m Region) Valid() (bool, error) {
	return husk.ValidateStruct(&m)
}
