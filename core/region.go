package core

import "github.com/louisevanderlith/husk"

type Region struct {
	Name      string
	StartChar string
	EndChar   string
	Countries []Country
}

func (m Region) Valid() (bool, error) {
	return husk.ValidateStruct(&m)
}

func GetRegionByCode(uniquevin string) (*Region, error) {
	record, err := ctx.Regions.FindFirst(byUniqueVIN(uniquevin))

	if err != nil {
		return nil, err
	}

	return record.Data().(*Region), nil
}
