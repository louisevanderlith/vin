package core

import (
	"github.com/louisevanderlith/husk/hsk"
)

type regionCalc func(result interface{}, obj Region) error

func (f regionCalc) Map(result interface{}, obj hsk.Record) error {
	return f(result, obj.GetValue().(Region))
}

func Manufacturers(year int) regionCalc {
	return func(result interface{}, obj Region) error {
		lst := *(result.(*map[string]struct{}))

		for _, country := range obj.Countries {
			for _, manufacturer := range country.Manufacturers {
				for _, plant := range manufacturer.AssemblyPlants {
					if year >= plant.StartYear && year <= plant.EndYear {
						lst[manufacturer.Name] = struct{}{}
					}
				}
			}
		}

		result = &lst

		return nil
	}
}

func Models(year int, manufacturerName string) regionCalc {
	return func(result interface{}, obj Region) error {
		lst := *(result.(*map[string]struct{}))

		for _, country := range obj.Countries {
			for _, manufacturer := range country.Manufacturers {
				if manufacturer.Name == manufacturerName {
					for _, plant := range manufacturer.AssemblyPlants {
						if year >= plant.StartYear && year <= plant.EndYear {
							for _, series := range plant.Series {
								if year >= series.StartYear && year <= series.EndYear {
									lst[series.Spec] = struct{}{}
								}
							}
						}
					}
				}
			}
		}

		result = &lst

		return nil
	}
}

func Trim(year int, manufacturer, model string) regionCalc {
	return func(result interface{}, obj Region) error {
		lst := *(result.(*map[string]struct{}))

		/*if obj.Year == year && obj.Series.Manufacturer == manufacturer && obj.Series.Model == model {
			lst[obj.Trim] = struct{}{}
		}*/

		result = &lst

		return nil
	}
}
