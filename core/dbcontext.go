package core

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/husk/collections"
	"github.com/louisevanderlith/husk/hsk"
	"github.com/louisevanderlith/husk/op"
	"github.com/louisevanderlith/husk/records"
	"os"
	"reflect"
	"strings"
)

type VINContext interface {
	CreateVIN(vin VIN) error
	ValidateVIN(fullvin string) error
	BuildInfo(fullvin string) (*VIN, error)
	FindVIN(fullvin string) (*VIN, error)
	GetVIN(key hsk.Key) (*VIN, error)
	GetAllVINS(page, size int) (records.Page, error)
}

func Context() VINContext {
	return ctx
}

func (c context) CreateVIN(m VIN) error {
	_, err := c.VIN.Create(m)
	return err
}

//ValidateVIN does exactly what it says. This is the first step in creating a VIN DB Entry.
func (c context) ValidateVIN(fullvin string) error {
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
func (c context) BuildInfo(fullvin string) (*VIN, error) {
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

func (c context) GetVIN(key hsk.Key) (*VIN, error) {
	rec, err := c.VIN.FindByKey(key)

	if err != nil {
		return nil, err
	}

	return rec.Data().(*VIN), nil
}

func (c context) FindVIN(fullvin string) (*VIN, error) {
	rec, err := c.VIN.FindFirst(byFullVIN(fullvin))

	if err != nil {
		return nil, err
	}

	return rec.Data().(*VIN), nil
}

func (c context) GetAllVINS(page, size int) (records.Page, error) {
	return c.VIN.Find(page, size, op.Everything())
}

type context struct {
	VIN     husk.Table
	Regions husk.Table
}

var ctx context

func CreateContext() {
	ctx = context{
		Regions: husk.NewTable(Region{}),
		VIN:     husk.NewTable(VIN{}),
	}

	seed()
}

func Shutdown() {
	ctx.Regions.Save()
	ctx.VIN.Save()
}

func seed() {
	regions, err := regionSeeds()

	if err != nil {
		panic(err)
	}

	err = ctx.Regions.Seed(regions)

	if err != nil {
		panic(err)
	}
}

func regionSeeds() (collections.Enumerable, error) {
	f, err := os.Open("db/regions.seed.json")

	if err != nil {
		return nil, err
	}

	var items []Region
	dec := json.NewDecoder(f)
	err = dec.Decode(&items)

	if err != nil {
		return nil, err
	}

	return collections.ReadOnlyList(reflect.ValueOf(items)), nil
}
