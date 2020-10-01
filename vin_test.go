package main

import (
	"github.com/louisevanderlith/vin/core"
	"log"
	"testing"
)

//Audi -- WAUZZZ8E88A025765
//Chev -- KL1MJ68036C084769
//Hyundai -- "KNHCU41DLCU177882
//Mustang -- 1ZVHT82H485113456
//Hyundai 2 --5NPEU46F77H259112
//Toyota -- JT152EEA100302159

//This is expectected from every test.
var expectations = core.VIN{
	Full:   "5NPEU46F77H259112",
	Unique: "5NPEU46F77H",
	Serial: 259112,
	WMInfo: core.WMInfo{
		Country:      "United States",
		Manufacturer: "Hyundai",
		VehicleType:  "PassengerCar",
		Region:       "North America",
	},
}

func init() {
	core.CreateContext()
}

func TestVIN_JustPrint(t *testing.T) {
	obj, err := core.Context().BuildInfo(expectations.Full)

	if err != nil {
		t.Error(err)
	}

	t.Log(obj)

	t.Fail()
}

func TestVIN_IsValid(t *testing.T) {
	in := "5NPEU46F77H259112"
	err := core.Context().ValidateVIN(in)

	if err != nil {
		t.Error(err)
	}
}

func TestVIN_NotValid(t *testing.T) {
	in := "5NBEU46F77H259112"
	err := core.Context().ValidateVIN(in)

	if err == nil {
		t.Error("Expecting error")
	}
}

func TestVIN_NoIOQ(t *testing.T) {
	in := "5NBEU46F77H259Q12"
	err := core.Context().ValidateVIN(in)

	if err == nil {
		t.Error("Expecting error")
	}
}

func TestDeconstruct_UniqueSerial_SerialCorrect(t *testing.T) {
	obj, err := core.Context().BuildInfo(expectations.Full)

	if err != nil {
		t.Error(err)
	}

	if obj.Serial != expectations.Serial {
		t.Errorf("expected %v, got %v", expectations.Serial, obj.Serial)
	}
}

func TestDeconstruct_UniqueSerial_UniqueCorrect(t *testing.T) {
	obj, err := core.Context().BuildInfo(expectations.Full)

	if err != nil {
		t.Error(err)
	}

	if obj.Unique != expectations.Unique {
		t.Errorf("expected %s, got %s", expectations.Unique, obj.Unique)
	}
}

func TestDeconstruct_WMI_ManufacturerCorrect(t *testing.T) {
	obj, err := core.Context().BuildInfo(expectations.Full)

	if err != nil {
		t.Error(err)
	}

	if obj.WMInfo.Manufacturer != expectations.WMInfo.Manufacturer {
		t.Errorf("expected %v, got %v", expectations.WMInfo.Manufacturer, obj.WMInfo.Manufacturer)
	}
}

func TestDeconstruct_WMI_CountryCorrect(t *testing.T) {
	obj, err := core.Context().BuildInfo(expectations.Full)

	if err != nil {
		t.Error(err)
	}

	if obj.WMInfo.Country != expectations.WMInfo.Country {
		t.Errorf("expected %v, got %v", expectations.WMInfo.Country, obj.WMInfo.Country)
	}
}

func TestDeconstruct_WMI_VehicleTypeCorrect(t *testing.T) {
	obj, err := core.Context().BuildInfo(expectations.Full)

	if err != nil {
		t.Error(err)
	}

	if obj.WMInfo.VehicleType != expectations.WMInfo.VehicleType {
		t.Errorf("expected %v, got %v", expectations.WMInfo.VehicleType, obj.WMInfo.VehicleType)
	}
}

func TestDeconstruct_WMI_RegionCorrect(t *testing.T) {
	obj, err := core.Context().BuildInfo(expectations.Full)

	if err != nil {
		t.Error(err)
	}

	if obj.WMInfo.Region != expectations.WMInfo.Region {
		t.Errorf("expected %v, got %v", expectations.WMInfo.Region, obj.WMInfo.Region)
	}
}

func TestDeconstruct_VDS_Toyota(t *testing.T) {
	obj, err := core.Context().BuildInfo("JT2MX83E2K0030681")

	if err != nil {
		t.Error(err)
	}

	log.Println(obj)
	t.Fail()
	//FINISH
}
