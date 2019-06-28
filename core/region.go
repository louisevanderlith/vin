package core

import (
	"strconv"

	"github.com/louisevanderlith/husk"
)

type Region struct {
	Name      string
	StartChar string
	EndChar   string
	Countries []Country
}

func (m Region) Valid() (bool, error) {
	return husk.ValidateStruct(&m)
}

func (r *Region) HasCode(regionCode string) bool {
	s := getCharWeight(r.StartChar)
	e := getCharWeight(r.EndChar)
	v := getCharWeight(regionCode)

	return s <= v && v <= e
}

func GetRegion(key husk.Key) (*Region, error) {
	rec, err := ctx.Regions.FindByKey(key)

	if err != nil {
		return nil, err
	}

	return rec.Data().(*Region), nil
}

func GetAllRegions(page, size int) husk.Collection {
	return ctx.Regions.Find(page, size, husk.Everything())
}

func GetRegionByCode(uniquevin string) (*Region, error) {
	record, err := ctx.Regions.FindFirst(byUniqueVIN(uniquevin))

	if err != nil {
		return nil, err
	}

	return record.Data().(*Region), nil
}

func getCharWeight(char string) int {
	if val, err := strconv.Atoi(char); err == nil {
		if val == 0 {
			return 36
		}

		return val + 26
	}

	//Alpha chars will return their index in the alphabet
	return int(char[0] % 32)
}

func (p Region) Update(key husk.Key) error {
	reg, err := ctx.Regions.FindByKey(key)

	if err != nil {
		return err
	}

	err = reg.Set(p)

	if err != nil {
		return err
	}

	defer ctx.Regions.Save()
	return ctx.Regions.Update(reg)
}
