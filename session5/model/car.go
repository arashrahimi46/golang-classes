package model

import "fmt"

type Car struct {
	Wheel          int
	Color          string
	Model          string
	Brand          string
	ProductionYear uint
	IsHachBack     bool
	ExtraOptions   interface{}
}

func NewCar(wheel int, color string,
	model string, brand string, productionYear uint, isHachBack bool) DriveAble {
	return &Car{
		Wheel:          wheel,
		Color:          color,
		Model:          model,
		Brand:          brand,
		ProductionYear: productionYear,
		IsHachBack:     isHachBack,
		ExtraOptions:   "ABS",
	}
}

func (c *Car) Drive(velocity int) int {
	fmt.Println("car is going with speed ", velocity)
	return velocity
}

func (c *Car) Park() bool {
	return false
}

func (c *Car) Break() bool {
	return true
}
