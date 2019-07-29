package vds

type VolvoVDS struct {
	Model string
}

func AnalyseVolvo(vds string, obj *VDSInfo) (interface{}, error) {
	tmp := new(VolvoVDS)

	macros := make(map[int]func(char string))
	macros[4] = tmp.Position4
	//macros[5] = tmp.Position5
	//6 - Series?
	//macros[7] = tmp.Position7
	//macros[8] = tmp.Position8

	for k, v := range vds {
		macro, ok := macros[k+4]

		if ok {
			macro(string(v))
		}
	}

	return tmp, nil
}

//Vehicle Type
func (v *VolvoVDS) Position4(char string) {
	switch char {
	case "G":
		v.Model = "S70, V70 BI-Fuel"
		break
	case "J":
		v.Model = "V70 BI-Fuel"
		break
	case "L":
		v.Model = "S70, V70, V70XC"
		break
	case "N":
		v.Model = "C70"
		break
	case "S":
		v.Model = "V70, V70XC, CX70"
		break
	}
}
