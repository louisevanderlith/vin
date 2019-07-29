package vds

type ToyotaVDS struct {
	Doors       int
	BodyStyle   string
	DriveTrain  string
	EngineModel string
	Safety      string
	Model       string
	Platform    string
}

func AnalyseToyota(vds string, obj *VDSInfo) (interface{}, error) {
	tmp := new(ToyotaVDS)

	macros := make(map[int]func(char string))
	macros[4] = tmp.Position4
	macros[5] = tmp.Position5
	//6 - Series?
	macros[7] = tmp.Position7
	macros[8] = tmp.Position8

	for k, v := range vds {
		macro, ok := macros[k+4]

		if ok {
			macro(string(v))
		}
	}

	return tmp, nil
}

//Body Type
func (v *ToyotaVDS) Position4(char string) {
	switch char {
	case "A":
		v.Doors = 2
		v.BodyStyle = "Sedan"
		v.DriveTrain = "2WD"
		//A/2DR sedan 2WD,
		break
	case "B":
		v.Doors = 4
		v.BodyStyle = "Sedan"
		v.DriveTrain = "2WD"
		//B/4DR sedan 2WD or 4DR truck 4WD,
		break
	case "C":
		v.Doors = 2
		v.BodyStyle = "Coupe"
		v.DriveTrain = "2WD"
		//C/2DR coupe 2WD,
		break
	case "D":
		v.Doors = 4
		v.BodyStyle = "Truck"
		v.DriveTrain = "4WD"
		//D/4DR truck 4WD,
		break
	case "E":
		v.Doors = 4
		v.BodyStyle = "Truck"
		v.DriveTrain = "2WD"
		//E/4DR truck 2WD,
		break
	case "G":
		v.Doors = 4
		v.BodyStyle = "Wagon"
		v.DriveTrain = "2WD"
		//G/4DR wagon 2WD,
		break
	case "H":
		v.Doors = 4
		v.BodyStyle = "Wagon"
		v.DriveTrain = "4WD"
		//H/4DR wagon 4WD,
		break
	case "K":
		v.Doors = 4
		v.BodyStyle = "Wagon"
		v.DriveTrain = "2WD"
		//K/4DR wagon 2WD,
		break
	case "L":
		v.Doors = 4
		v.BodyStyle = "Wagon"
		v.DriveTrain = "4WD"
		//L/4DR wagon 4WD or 4DR truck 4WD,
		break
	case "M":
		v.Doors = 5
		v.BodyStyle = "Van"
		v.DriveTrain = "2WD"
		//M/5DR van 2WD,
		break
	case "N":
		v.Doors = 2
		v.BodyStyle = "Regular Cab"
		v.DriveTrain = "2WD"
		//N/2DR regular cab truck 2WD,
		break
	case "P":
		v.Doors = 2
		v.BodyStyle = "Regular Cab"
		v.DriveTrain = "4WD"
		//P/2DR regular cab truck 4WD,
		break
	case "S":
		v.Doors = 3
		v.BodyStyle = "Extended Cab"
		v.DriveTrain = "4WD"
		//S/3DR liftback 4WD,
		break
	case "T":
		v.Doors = 2
		v.BodyStyle = "Extended Cab"
		v.DriveTrain = "2WD"
		//T/2DR extended cab truck 2WD,
		break
	case "X":
		v.Doors = 5
		v.BodyStyle = "SUV"
		//X/5DR SUV,
		break
	case "W":
		v.Doors = 2
		v.BodyStyle = "Extended Cab"
		v.DriveTrain = "4WD"
		//W/2DR extended cab 4WD,
		break
	case "Y":
		v.BodyStyle = "Sport Van"
		//Y/sport van and
		break
	case "Z":
		v.Doors = 5
		v.BodyStyle = "Wagon"
		v.DriveTrain = "2WD"
		//Z/5DR wagon 2WD.
		break
	}
}

//Engine
func (v *ToyotaVDS) Position5(char string) {
	switch char {
	case "4":
		v.EngineModel = "7A-FE"
		//4/7A-FE Lean Burn;
		break
	case "A":
		v.EngineModel = "3MZ-FE"
		//A/3MZ-FE;
		break
	case "B":
		v.EngineModel = "1NZ-FXE"
		//B/1NZ-FXE or Toyota AZ engine#2AZ-FXE|2AZ-FXE;
		break
	case "D":
		v.EngineModel = "2ZZ-GE"
		//D/2ZZ-GE;
		break
	case "E":
		v.EngineModel = "2AZ-FE"
		//E/2AZ-FE;
		break
	case "F":
		v.EngineModel = "1MZ-FE"
		//F/1MZ-FE or 2AR-FE;
		break
	case "G":
		v.EngineModel = "5S-FE"
		//G/5S-FE;
		break
	case "H":
		v.EngineModel = "1AZ-FE"
		//H/1AZ-FE;
		break
	case "J":
		v.EngineModel = "1FZ-FE"
		//J/1FZ-FE;
		break
	case "K":
		v.EngineModel = "2GR-FE"
		//K/2GR-FE;
		break
	case "L":
		v.EngineModel = "2RZ-FE"
		//L/2RZ-FE;
		break
	case "M":
		v.EngineModel = "3RZ-FE"
		//M/3RZ-FE;
		break
	case "N":
		v.EngineModel = "5VZ-FE"
		//N/5VZ-FE or	  2ZR-FXE;
		break
	case "P":
		v.EngineModel = "3S-FE"
		//P/3S-FE;
		break
	case "R":
		v.EngineModel = "1ZZ-FE"
		//R/1ZZ-FE;
		break
	case "S":
		v.EngineModel = "1BM"
		//S/1BM or Electric;
		break
	case "T":
		v.EngineModel = "3S-GTE"
		//T/3S-GTE;
		break
	case "U":
		v.EngineModel = "1GR-FE"
		//U/1GR-FE or 2ZR-FE;
		break
	case "V":
		v.EngineModel = "1NR-FE"
		//V/1NR-FE and
		break
	case "Y":
		v.EngineModel = "3UR-FE"
		//Y/3UR-FE.
		break
	}
}

//Series...
func (v *ToyotaVDS) Position6(char string) {

}

//Safety Features
func (v *ToyotaVDS) Position7(char string) {
	switch char {
	case "0":
		v.Safety = "Manual Belts with 2 Airbags and Curtain Airbags"
		break
	case "1":
		v.Safety = "Manual Belt"
		break
	case "2":
		v.Safety = "Manual Belts with Driver Side Airbag"
		break
	case "3":
		v.Safety = "Manual Belts with 2 Airbags"
		break
	case "6":
		v.Safety = "Manual Belts with 2 Airbags, Side Airbags, Curtain Airbags and Knee Airbag for driver"
		break
	case "7":
		v.Safety = "Manual Belts with 2 Airbags and Knee Airbag for driver"
		break
	case "8":
		v.Safety = "Manual Belts with 2 Airbags and Side Airbags"
		break
	case "D":
		v.Safety = "Manual Belts with 2 Airbags, Side Airbags, Three-Row Curtain Airbags and Knee Airbag"
		break
	case "F":
		v.Safety = "Manual Belts with 2 Airbags, Side Airbags and Knee Airbag"
		break
	}
}

func (v *ToyotaVDS) Position8(char string) {
	switch char {
	case "0":
		v.Model = "MR2 Spyder"
		break
	case "1":
		v.Model = "Tundra"
		break
	case "3":
		v.Model = "Yaris"
		break
	case "4":
		v.Model = "xB"
		break
	case "7":
		v.Model = "tC"
		break
	case "A":
		v.Model = "Highlander, Sequoia, Celica and Supra"
		break
	case "B":
		v.Model = "Avalon"
		break
	case "C":
		v.Model = "Sienna and Previa"
		break
	case "D":
		v.Model = "T100"
	case "E":
		v.Model = "Corolla or Matrix"
		break
	case "F":
		v.Model = "FJ Cruiser"
		break
	case "G":
		v.Model = "Hilux"
		break
	case "H":
		v.Model = "Highlander"
		break
	case "J":
		v.Model = "Land Cruiser"
		break
	case "K":
		v.Model = "Camry"
		break
	case "L":
		v.Model = "Tercel and Paseo"
		break
	case "M":
		v.Model = "Previa"
		break
	case "N":
		v.Model = "Tacoma and older trucks"
		break
	case "P":
		v.Model = "Camry Solara"
		break
	case "R":
		v.Model = "4Runner and Corolla"
		break
	case "T":
		v.Model = "Celica"
		v.Platform = "FWD"
		break
	case "U":
		v.Model = "Prius"
		break
	case "V":
		v.Model = "RAV4"
		break
	case "W":
		v.Model = "MR2 non Spyder"
		break
	case "X":
		v.Model = "Cressida"
		break
	}
}
