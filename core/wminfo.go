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

	regionCode := uniquevin[:1]
	countryCode := uniquevin[1:2]

	for i := 0; i < len(region.Countries); i++ {
		country := region.Countries[i]

		if country.RegionCode == regionCode && country.HasCode(countryCode) {
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
