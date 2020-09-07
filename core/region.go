package core

import (
	"github.com/louisevanderlith/husk/hsk"
	"github.com/louisevanderlith/husk/op"
	"github.com/louisevanderlith/husk/records"
	"github.com/louisevanderlith/husk/validation"
	"strconv"
)

type Region struct {
	Name      string
	StartChar string
	EndChar   string
	Countries []Country
}

func (m Region) Valid() error {
	return validation.Struct(m)
}

func (r Region) HasCode(regionCode string) bool {
	s := getCharWeight(r.StartChar)
	e := getCharWeight(r.EndChar)
	v := getCharWeight(regionCode)

	return s <= v && v <= e
}

func GetRegion(key hsk.Key) (Region, error) {
	rec, err := ctx.Regions.FindByKey(key)

	if err != nil {
		return Region{}, err
	}

	return rec.Data().(Region), nil
}

func GetAllRegions(page, size int) (records.Page, error) {
	return ctx.Regions.Find(page, size, op.Everything())
}

func GetRegionByCode(uniquevin string) (Region, error) {
	record, err := ctx.Regions.FindFirst(byUniqueVIN(uniquevin))

	if err != nil {
		return Region{}, err
	}

	return record.Data().(Region), nil
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

func (p Region) Update(key hsk.Key) error {
	return ctx.Regions.Update(key, p)
}
