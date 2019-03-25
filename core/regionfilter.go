package core

import (
	"github.com/louisevanderlith/husk"
)

type regionFilter func(obj *Region) bool

func (f regionFilter) Filter(obj husk.Dataer) bool {
	return f(obj.(*Region))
}

func byUniqueVIN(uniquevin string) regionFilter {
	regionChar := uniquevin[:1]

	return func(obj *Region) bool {
		return obj.HasCode(regionChar)
	}
}
