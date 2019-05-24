package core

import (
	"github.com/louisevanderlith/husk"
)

type vinFilter func(obj *VIN) bool

func (f vinFilter) Filter(obj husk.Dataer) bool {
	return f(obj.(*VIN))
}

func byFullVIN(full string) vinFilter {
	return func(obj *VIN) bool {
		return obj.Full == full
	}
}
