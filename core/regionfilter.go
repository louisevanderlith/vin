package core

import (
	"github.com/louisevanderlith/husk/hsk"
)

type regionFilter func(obj Region) bool

func (f regionFilter) Filter(obj hsk.Record) bool {
	return f(obj.Data().(Region))
}

func byUniqueVIN(uniquevin string) regionFilter {
	regionChar := uniquevin[:1]

	return func(obj Region) bool {
		return obj.HasCode(regionChar)
	}
}
