package core

import "github.com/louisevanderlith/husk"

type vehicleFilter func(obj *Vehicle) bool

func (f vehicleFilter) Filter(obj husk.Dataer) bool {
	return f(obj.(*Vehicle))
}

func byFullVIN(fullvin string) vehicleFilter {
	return func(obj *Vehicle) bool {
		return obj.VIN.Full == fullvin
	}
}
