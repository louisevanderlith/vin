package core

import "testing"

func TestVIN_IsValid(t *testing.T) {
	in := "5NPEU46F77H259112"
	err := ValidateVIN(in)

	if err != nil {
		t.Error(err)
	}
}

func TestVIN_NotValid(t *testing.T) {
	in := "5NBEU46F77H259112"
	err := ValidateVIN(in)

	if err == nil {
		t.Error("Expecting error")
	}
}

func TestVIN_NoIOQ(t *testing.T) {
	in := "5NBEU46F77H259Q12"
	err := ValidateVIN(in)

	if err == nil {
		t.Error("Expecting error")
	}
}

func TestDeconstruct_UniqueSerial_SerialCorrect(t *testing.T) {
	expect := VIN{
		Full:   "5NPEU46F77H259112",
		Serial: 259112,
	}

	obj := newVIN("5NPEU46F77H259112")

	err := obj.deconstruct()

	if err != nil {
		t.Error(err)
	}

	if obj.Serial != expect.Serial {
		t.Errorf("expected %v, got %v", expect.Serial, obj.Serial)
	}
}

func TestDeconstruct_UniqueSerial_UniqueCorrect(t *testing.T) {
	expect := VIN{
		Full:   "5NPEU46F77H259112",
		Unique: "5NPEU46F77H",
	}

	obj := newVIN("5NPEU46F77H259112")

	err := obj.deconstruct()

	if err != nil {
		t.Error(err)
	}

	if obj.Unique != expect.Unique {
		t.Errorf("expected %s, got %s", expect.Unique, obj.Unique)
	}
}
