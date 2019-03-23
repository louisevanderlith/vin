package core

type WMInfo struct {
	Region       string
	Country      string
	Manufacturer string
	VehicleType  VehicleType
}

func FindWMInfo(uniquevin string) (WMInfo, error) {
	result := WMInfo{}

	region, err := GetRegionByCode(uniquevin)

	if err != nil {
		return result, err
	}

	result.Region = region.Name

	countryCode := uniquevin[1:1]
	for i := 0; i < len(region.Countries); i++ {
		country := region.Countries[i]

		if country.StartChar <= countryCode && country.EndChar >= countryCode {
			result.Country = country.Name
			wmi := uniquevin[:3]

			for j := 0; j < len(country.Manufacturers); j++ {
				manufacturer := country.Manufacturers[j]

				if manufacturer.WMICode == wmi {
					result.Manufacturer = manufacturer.Name
					result.VehicleType = manufacturer.VehicleType

					break
				}
			}

			break
		}
	}

	return result, nil
}
