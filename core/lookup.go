package core

func GetManufacturers(year int) (map[string]struct{}, error) {
	result := make(map[string]struct{})
	err := ctx.Regions.Calculate(&result, Manufacturers(year))

	if err != nil {
		return nil, err
	}

	return result, nil
}

func GetModels(year int, manufacturer string) (map[string]struct{}, error) {
	result := make(map[string]struct{})
	err := ctx.Regions.Calculate(&result, Models(year, manufacturer))

	if err != nil {
		return nil, err
	}

	return result, nil
}

func GetTrims(year int, manufacturer, model string) (map[string]struct{}, error) {
	result := make(map[string]struct{})
	err := ctx.Regions.Calculate(&result, Trim(year, manufacturer, model))

	if err != nil {
		return nil, err
	}

	return result, nil
}
