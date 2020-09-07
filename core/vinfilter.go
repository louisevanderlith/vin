package core

import (
	"github.com/louisevanderlith/husk/hsk"
)

type vinFilter func(obj VIN) bool

func (f vinFilter) Filter(obj hsk.Record) bool {
	return f(obj.Data().(VIN))
}

func byFullVIN(full string) vinFilter {
	return func(obj VIN) bool {
		return obj.Full == full
	}
}
