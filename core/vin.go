package core

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/louisevanderlith/husk"
)

//VIN is the key to the entire vehicle database.
type VIN struct {
	Full   string `hsk:"size(17)"`
	Unique string `hsk:"min(2)"`
	Serial int
	WMInfo WMInfo
}

func newVIN(fullvin string) (*VIN, error) {
	vin := &VIN{
		Full: fullvin,
	}

	err := vin.deconstruct()

	if err != nil {
		return nil, err
	}

	return vin, nil
}

//Valid checks if the object's values meets the data requirements
func (m VIN) Valid() (bool, error) {
	return husk.ValidateStruct(&m)
}

func (m VIN) Create() (husk.Recorder, error) {
	item, err := ctx.VIN.FindFirst(byFullVIN(m.Full))

	//If Found, just return the record
	if err == nil {
		return item, nil
	}

	cset := ctx.VIN.Create(m)

	if cset.Error != nil {
		return nil, cset.Error
	}

	return cset.Record, nil
}

//ValidateVIN does exactly what it says. This is the first step in creating a VIN DB Entry.
func ValidateVIN(fullvin string) error {
	if len(fullvin) != 17 {
		return errors.New("not correct length")
	}

	if strings.ContainsAny(fullvin, "IOQ") {
		return errors.New("found illegal characters")
	}

	checkDigit := fullvin[8:9]
	score := calculateScore(fullvin)

	if checkDigit != score {
		return fmt.Errorf("check digit %s is invalid for %s", checkDigit, score)
	}

	return nil
}

//BuildInfo tries to extract information from VIN number
func BuildInfo(fullvin string) (*VIN, error) {
	vin, err := newVIN(fullvin)

	if err != nil {
		return nil, err
	}

	err = vin.deconstruct()

	if err != nil {
		return nil, err
	}

	return vin, nil
}

//deconstruct will attempt to populat as much detail as possible for the given VIN
func (m *VIN) deconstruct() error {
	m.Unique, m.Serial = getUniqueSerial(m.Full)
	wmiInfo, err := FindWMInfo(m.Unique)

	if err != nil {
		return err
	}

	m.WMInfo = wmiInfo

	return nil
}

func getUniqueSerial(fullvin string) (string, int) {
	serial, _ := strconv.Atoi(fullvin[11:])
	return fullvin[:11], serial
}

func calculateScore(fullvin string) string {
	result := 0

	digitMap := getCharacterMap()
	weights := []int{8, 7, 6, 5, 4, 3, 2, 10, 0, 9, 8, 7, 6, 5, 4, 3, 2}

	for k, v := range fullvin {
		value := 0
		strVal := string(v)
		value, ok := digitMap[strVal]

		//If the character is not found, it's a digit.
		if !ok {
			val, err := strconv.Atoi(strVal)

			if err != nil {
				panic(err)
			}

			value = val
		}

		result += value * weights[k]
	}

	mod := result % 11

	if mod == 10 {
		return "X"
	}

	return strconv.Itoa(mod)
}

func getCharacterMap() map[string]int {
	digitMap := make(map[string]int)
	digitMap["A"] = 1
	digitMap["B"] = 2
	digitMap["C"] = 3
	digitMap["D"] = 4
	digitMap["E"] = 5
	digitMap["F"] = 6
	digitMap["G"] = 7
	digitMap["H"] = 8
	digitMap["J"] = 1
	digitMap["K"] = 2
	digitMap["L"] = 3
	digitMap["M"] = 4
	digitMap["N"] = 5
	digitMap["P"] = 7
	digitMap["R"] = 9
	digitMap["S"] = 2
	digitMap["T"] = 3
	digitMap["U"] = 4
	digitMap["V"] = 5
	digitMap["W"] = 6
	digitMap["X"] = 7
	digitMap["Y"] = 8
	digitMap["Z"] = 9

	return digitMap
}

/*
func doesVINExist(fullvin string) (husk.Recorder, bool) {
	result, err := ctx.Vehicles.FindFirst(byFullVIN(fullvin))

	if err != nil {
		return nil, false
	}

	return result, true
}
*/
