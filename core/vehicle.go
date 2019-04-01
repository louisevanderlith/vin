package core

import "github.com/louisevanderlith/husk"

type Vehicle struct {
	VIN            *VIN
	Series         SeriesInfo
	Colour         string
	Year           int
	AssemblyPlant  *AssemblyPlant
	EngineHistory  []*EngineHistory
	GearboxHistory []*GearboxHistory
	Alterations    []*Alteration
}

type SeriesInfo struct {
	Model        string
	Manufacturer string
	Etc          string
}

func (m Vehicle) Valid() (bool, error) {
	return husk.ValidateStruct(&m)
}

//BuildVehicle will create a full as possible detail record
func BuildVehicle(fullvin string) (husk.Recorder, error) {
	result, ok := doesVINExist(fullvin)

	if ok {
		return result, nil
	}

	vin, err := newVIN(fullvin)

	if err != nil {
		return nil, err
	}

	obj := Vehicle{
		VIN: vin,
		//Series:
	}

	set := ctx.Vehicles.Create(obj)

	if set.Error != nil {
		return nil, set.Error
	}

	ctx.Vehicles.Save()

	return set.Record, nil
}
