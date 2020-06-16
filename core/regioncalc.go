package core

import "github.com/louisevanderlith/husk"

type regionCalc func(result interface{}, obj Region) error

func (f regionCalc) Calc(result interface{}, obj husk.Dataer) error {
	return f(result, obj.(Region))
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

func Models(year int, manufacturer string) regionCalc {
	return func(result interface{}, obj Region) error {
		lst := *(result.(*map[string]struct{}))

		/*for _, country := range obj.Countries {
			for _, manufacturer := range country.Manufacturers {
				for _, plant := range manufacturer.AssemblyPlants {
					if year >= plant.StartYear && year <= plant.EndYear {
						for _, series := range plant.Series {
							series.
						}

						lst[plant.] = struct{}{}
					}
				}
			}
		}


		if obj.Year == year && obj.Series.Manufacturer == manufacturer {
			lst[obj.Series.Model] = struct{}{}
		}*/

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
